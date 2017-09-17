package k8s

import "k8s.io/client-go/kubernetes"

type Service interface {
	DeploymentsGetter
	PodsGetter
}

type service struct {
	clientSet *kubernetes.Clientset
}

func NewService() (Service, error) {
	client, err := NewInClusterClient()
	if err != nil {
		return nil, err
	}
	return service{
		clientSet: client,
	}, nil
}

func NewServiceFromConfig(path *string) (Service, error) {
	client, err := NewClientFromConfig(path)
	if err != nil {
		return nil, err
	}
	return service{
		clientSet: client,
	}, nil
}

func (s service) Deployments(namespace string) Deployments {
	return s.Deployments(namespace)
}

func (s service) Pods(namespace string) Pods {
	return s.Pods(namespace)
}
