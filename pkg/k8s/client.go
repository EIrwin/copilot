package k8s

import (
	"errors"
	"log"

	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func NewClient(configPath string) (*kubernetes.Clientset, error) {
	if configPath == "" {
		return NewClientFromConfig(&configPath)
	}
	return NewInClusterClient()
}

func NewInClusterClient() (*kubernetes.Clientset, error) {

	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return clientset, nil
}

func NewClientFromConfig(configPath *string) (*kubernetes.Clientset, error) {
	var config *rest.Config
	config, err := clientcmd.BuildConfigFromFlags("", *configPath)
	if err != nil {
		return nil, err
	}
	return createClientSet(config), nil
}

func createClientSet(c *rest.Config) *kubernetes.Clientset {
	clientset, err := kubernetes.NewForConfig(c)
	if err != nil {
		log.Panicln(errors.New(err.Error()))
	}
	return clientset
}
