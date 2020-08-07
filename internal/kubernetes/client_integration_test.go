package kubernetes

import (
	"github.com/stretchr/testify/suite"
	"k8s.io/client-go/kubernetes"
	"os"
	"testing"
)

type IntegrationTestSuite struct {
	suite.Suite
	testClient KubernetesClient
	clientSet  kubernetes.Interface
}

func TestIntegrationTestSuite(t *testing.T) {
	value, available := os.LookupEnv("ENABLE_INTEGRATION_TEST")
	if available == true && value == "true" {
		suite.Run(t, new(IntegrationTestSuite))
	}
}
