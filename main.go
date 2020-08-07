package main

import (
	"bytes"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"crud-toy/config"
	"crud-toy/internal/kubernetes"
)

func main() {
	var proctorConfig = config.ProctorConfig{
		KubeConfig:       "out-of-cluster",
		KubeContext:      "minikube",
		DefaultNamespace: "default",
	}

	var newClient, err = kubernetes.NewKubernetesClient(proctorConfig)

	if err != nil {
		fmt.Println("ERROR! Unable to create new clientSet: ", err)
	} else {
		fmt.Printf("new Kube client created: %+v\n", newClient)
	}

	listOptions := v1.ListOptions{
		LabelSelector: "component=etcd",
	}

	podList, err := newClient.ListPod("kube-system", listOptions)
	if err != nil {
		panic(err)
	}

	readCloser, err := newClient.GetPodLogs(&podList[0])
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(readCloser)

	defer readCloser.Close()

	newStr := buf.String()

	go func() {
		fmt.Println(newStr)
	}()

    // Program Terminated
}
