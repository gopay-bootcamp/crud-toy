package kubernetes

import (
	"context"
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"io"
	batch "k8s.io/api/batch/v1"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
	"time"
	"crud-toy/config"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	kubeRestClient "k8s.io/client-go/rest"

)

var (
	typeMeta     meta.TypeMeta
	namespace    string
	timeoutError = errors.New("timeout when waiting job to be available")

	ctx    context.Context
	cancel context.CancelFunc
)

func init() {
	typeMeta = meta.TypeMeta{
		Kind:       "Job",
		APIVersion: "batch/v1",
	}
	namespace = "default"
}

type kubernetesClient struct {
	clientSet kubernetes.Interface
}

type KubernetesClient interface {
	ExecuteJobWithCommand(imageName string, args map[string]string, commands []string) (string, error)
	ExecuteJob(imageName string, args map[string]string) (string, error)
	JobExecutionStatus(jobName string) (string, error)
	GetPodLogs(pod *v1.Pod) (io.ReadCloser, error)
	WaitForReadyJob(executionName string, waitTime time.Duration) error
	WaitForReadyPod(executionName string, waitTime time.Duration) (*v1.Pod, error)
	ListPod(namespace string, options meta.ListOptions) ([]v1.Pod, error)
}

func (client *kubernetesClient) ListPod(namespace string, options meta.ListOptions) ([]v1.Pod, error) {
	coreV1 := client.clientSet.CoreV1()
	pods := coreV1.Pods(namespace)

	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()
	podList, err := pods.List(ctx, options)
	if err != nil {
		return nil, err
	}

	return podList.Items, nil
}

func NewClientSet(proctorConfig config.ProctorConfig) (*kubernetes.Clientset, error) {
	var kubeConfig *kubeRestClient.Config

	if proctorConfig.KubeConfig == "out-of-cluster" {
		//logger.Info("service is running outside kube cluster")
		fmt.Println("service is running outside kube cluster")
		home := os.Getenv("HOME")

		kubeConfigPath := filepath.Join(home, ".kube", "config")

		configOverrides := &clientcmd.ConfigOverrides{}
		if config.Config().KubeContext != "default" {
			configOverrides.CurrentContext = config.Config().KubeContext
		}

		var err error
		kubeConfig, err = clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
			&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeConfigPath},
			configOverrides).ClientConfig()
		if err != nil {
			return nil, err
		}

	} else {
		var err error
		kubeConfig, err = kubeRestClient.InClusterConfig()
		if err != nil {
			return nil, err
		}
	}

	clientSet, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		return nil, err
	}
	return clientSet, nil
}

func NewKubernetesClient(proctorConfig config.ProctorConfig) (KubernetesClient, error) {
	client := &kubernetesClient{}

	var err error
	client.clientSet, err = NewClientSet(proctorConfig)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func uniqueName() string {
	return "proctor" + "-" + uuid.NewV4().String()
}

func jobLabel(executionName string) map[string]string {
	return map[string]string{
		"job": executionName,
	}
}

func jobLabelSelector(executionName string) string {
	return fmt.Sprintf("job=%s", executionName)
}

func getEnvVars(envMap map[string]string) []v1.EnvVar {
	var envVars []v1.EnvVar
	for k, v := range envMap {
		envVar := v1.EnvVar{
			Name:  k,
			Value: v,
		}
		envVars = append(envVars, envVar)
	}
	return envVars
}

func watcherError(resource string, listOptions meta.ListOptions) error {
	return fmt.Errorf("watch error when waiting for %s with list option %v", resource, listOptions)
}

func (client *kubernetesClient) ExecuteJob(imageName string, envMap map[string]string) (string, error) {
	return client.ExecuteJobWithCommand(imageName, envMap, []string{})
}

func (client *kubernetesClient) ExecuteJobWithCommand(imageName string, envMap map[string]string, command []string) (string, error) {
	executionName := uniqueName()
	label := jobLabel(executionName)

	batchV1 := client.clientSet.BatchV1()
	kubernetesJobs := batchV1.Jobs(namespace)

	container := v1.Container{
		Name:  executionName,
		Image: imageName,
		Env:   getEnvVars(envMap),
	}

	if len(command) != 0 {
		container.Command = command
	}

	podSpec := v1.PodSpec{
		Containers:         []v1.Container{container},
		RestartPolicy:      v1.RestartPolicyNever,
		ServiceAccountName: config.Config().KubeServiceAccountName,
	}

	objectMeta := meta.ObjectMeta{
		Name:        executionName,
		Labels:      label,
		Annotations: config.Config().JobPodAnnotations,
	}

	template := v1.PodTemplateSpec{
		ObjectMeta: objectMeta,
		Spec:       podSpec,
	}

	jobSpec := batch.JobSpec{
		Template:              template,
		ActiveDeadlineSeconds: config.Config().KubeJobActiveDeadlineSeconds,
		BackoffLimit:          config.Config().KubeJobRetries,
	}

	jobToRun := batch.Job{
		TypeMeta:   typeMeta,
		ObjectMeta: objectMeta,
		Spec:       jobSpec,
	}

	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	_, err := kubernetesJobs.Create(ctx, &jobToRun, meta.CreateOptions{})
	if err != nil {
		return "", err
	}
	return executionName, nil
}

func (client *kubernetesClient) JobExecutionStatus(executionName string) (string, error) {
	batchV1 := client.clientSet.BatchV1()
	kubernetesJobs := batchV1.Jobs(namespace)
	listOptions := meta.ListOptions{
		TypeMeta:      typeMeta,
		LabelSelector: jobLabelSelector(executionName),
	}

	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()
	watchJob, err := kubernetesJobs.Watch(ctx, listOptions)
	if err != nil {
		return "FAILED", err
	}

	resultChan := watchJob.ResultChan()
	defer watchJob.Stop()
	var event watch.Event
	var jobEvent *batch.Job

	for event = range resultChan {
		if event.Type == watch.Error {
			return "JOB_EXECUTION_STATUS_FETCH_ERROR", nil
		}

		jobEvent = event.Object.(*batch.Job)
		if jobEvent.Status.Succeeded >= int32(1) {
			return "SUCCEEDED", nil
		} else if jobEvent.Status.Failed >= int32(1) {
			return "FAILED", nil
		}
	}

	return "NO_DEFINITIVE_JOB_EXECUTION_STATUS_FOUND", nil
}

func (client *kubernetesClient) GetPodLogs(pod *v1.Pod) (io.ReadCloser, error) {
	//logger.Debug("reading pod logs for: ", pod.Name)
	podLogOpts := v1.PodLogOptions{
		Follow: true,
	}

	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	request := client.clientSet.CoreV1().Pods(pod.Namespace).GetLogs(pod.Name, &podLogOpts)
	response, err := request.Stream(ctx)

	if err != nil {
		return nil, err
	}
	return response, nil
}

func (client *kubernetesClient) WaitForReadyPod(executionName string, waitTime time.Duration) (*v1.Pod, error) {
	coreV1 := client.clientSet.CoreV1()
	kubernetesPods := coreV1.Pods(namespace)
	listOptions := meta.ListOptions{
		LabelSelector: jobLabelSelector(executionName),
	}

	var err error

	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	for i := 0; i < config.Config().KubeWaitForResourcePollCount; i += 1 {
		watchJob, watchErr := kubernetesPods.Watch(ctx, listOptions)
		if watchErr != nil {
			err = watchErr
			continue
		}

		timeoutChan := time.After(waitTime)
		resultChan := watchJob.ResultChan()

		var pod *v1.Pod
		for {
			select {
			case event := <-resultChan:
				if event.Type == watch.Error {
					err = watcherError("pod", listOptions)
					watchJob.Stop()
					break
				}

				// Ignore empty events
				if event.Object == nil {
					continue
				}

				pod = event.Object.(*v1.Pod)
				if pod.Status.Phase == v1.PodRunning || pod.Status.Phase == v1.PodSucceeded || pod.Status.Phase == v1.PodFailed {
					watchJob.Stop()
					return pod, nil
				}
			case <-timeoutChan:
				err = timeoutError
				watchJob.Stop()
				break
			}
			if err != nil {
				watchJob.Stop()
				break
			}
		}
	}

	//logger.Info("Wait for ready pod return pod ", nil, " and error ", err)
	fmt.Println("Wait for ready pod return pod ", nil, " and error ", err)
	return nil, err
}

func (client *kubernetesClient) WaitForReadyJob(executionName string, waitTime time.Duration) error {
	batchV1 := client.clientSet.BatchV1()
	jobs := batchV1.Jobs(namespace)
	listOptions := meta.ListOptions{
		TypeMeta:      typeMeta,
		LabelSelector: jobLabelSelector(executionName),
	}

	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	var err error
	for i := 0; i < config.Config().KubeWaitForResourcePollCount; i += 1 {
		watchJob, watchErr := jobs.Watch(ctx, listOptions)
		if watchErr != nil {
			err = watchErr
			continue
		}

		timeoutChan := time.After(waitTime)
		resultChan := watchJob.ResultChan()

		var job *batch.Job
		for {
			select {
			case event := <-resultChan:
				if event.Type == watch.Error {
					err = watcherError("job", listOptions)
					break
				}

				// Ignore empty events
				if event.Object == nil {
					continue
				}

				job = event.Object.(*batch.Job)
				if job.Status.Active >= 1 || job.Status.Succeeded >= 1 || job.Status.Failed >= 1 {
					watchJob.Stop()
					return nil
				}
			case <-timeoutChan:
				err = timeoutError
				break
			}
			if err != nil {
				watchJob.Stop()
				break
			}
		}
	}

	return err
}
