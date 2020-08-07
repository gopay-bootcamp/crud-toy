package kubernetes

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	v1 "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	fakeclientset "k8s.io/client-go/kubernetes/fake"
	batch "k8s.io/client-go/kubernetes/typed/batch/v1"
	"net/http"
	"os"
	"crud-toy/config"
	"testing"
)


type ClientTestSuite struct {
	suite.Suite
	testClient 				KubernetesClient
	testKubernetesJobs		batch.JobInterface
	fakeClientSet			*fakeclientset.Clientset
	jobName 				string
	podName 				string
	fakeClientSetStreaming 	*fakeclientset.Clientset
	fakeHTTPClient			*http.Client
	testClientStreaming		KubernetesClient
}

func (suite *ClientTestSuite) SetupTest() {
	suite.fakeClientSet = fakeclientset.NewSimpleClientset()
	suite.testClient = &kubernetesClient{
		clientSet: suite.fakeClientSet,
	}
	suite.jobName = "job1"
	suite.podName = "pod1"
	namespace := config.Config().DefaultNamespace
	suite.fakeClientSetStreaming = fakeclientset.NewSimpleClientset(&v1.Pod{
		TypeMeta: meta.TypeMeta{
			Kind: "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: meta.ObjectMeta{
			Name: suite.podName,
			Namespace: namespace,
			Labels: map[string]string{
				"tag": "",
				"job": suite.jobName,
			},
		},
		Status: v1.PodStatus{
			Phase: v1.PodSucceeded,
		},
	})

	suite.fakeHTTPClient =&http.Client{}
	suite.testClientStreaming = &kubernetesClient{
		clientSet: suite.fakeClientSetStreaming,
	}
}

func (suite *ClientTestSuite) TestJobExecution() {
	t := suite.T()
	_ = os.Setenv("PROCTOR_JOB_POD_ANNOTATIONS", "{\"key.one\":\"true\"}")
	_ = os.Setenv("PROCTOR_KUBE_SERVICE_ACCOUNT_NAME", "default")
	config.Reset()
	envVarsForContainer := map[string]string{"SAMPLE_ARG": "sample-value"}
	sampleImageName := "img1"

	executedJobName, err := suite.testClient.ExecuteJob(sampleImageName, envVarsForContainer)
	assert.NoError(t, err)

	typeMeta := meta.TypeMeta{
		Kind: "Job",
		APIVersion: "batch/v1",
	}

	listOptions := meta.ListOptions{
		TypeMeta: typeMeta,
		LabelSelector: jobLabelSelector(executedJobName),
	}
	namespace := config.Config().DefaultNamespace
	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()
	listOfJobs, err := suite.fakeClientSet.BatchV1().Jobs(namespace).List(ctx, listOptions)
	assert.NoError(t, err)
	executedJob := listOfJobs.Items[0]

	assert.Equal(t, typeMeta, executedJob.TypeMeta)

	assert.Equal(t, executedJobName, executedJob.ObjectMeta.Name)
	assert.Equal(t, executedJobName, executedJob.Spec.Template.ObjectMeta.Name)

	expectedLabel := jobLabel(executedJobName)
	assert.Equal(t, expectedLabel, executedJob.ObjectMeta.Labels)
	assert.Equal(t, expectedLabel, executedJob.Spec.Template.ObjectMeta.Labels)
	assert.Equal(t, map[string]string{"key.one": "true"}, executedJob.Spec.Template.Annotations)
	assert.Equal(t, "default", executedJob.Spec.Template.Spec.ServiceAccountName)

	assert.Equal(t, config.Config().KubeJobActiveDeadlineSeconds, executedJob.Spec.ActiveDeadlineSeconds)
	assert.Equal(t, config.Config().KubeJobRetries, executedJob.Spec.BackoffLimit)

	assert.Equal(t, v1.RestartPolicyNever, executedJob.Spec.Template.Spec.RestartPolicy)

	container := executedJob.Spec.Template.Spec.Containers[0]
	assert.Equal(t, executedJobName, container.Name)

	assert.Equal(t, sampleImageName, container.Image)

	expectedEnvVars := getEnvVars(envVarsForContainer)
	assert.Equal(t, expectedEnvVars, container.Env)
}

func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}