package k8s

import "k8s.io/client-go/kubernetes"

type Service interface {
	DeploymentsGetter
	PodsGetter
	ServicesGetter
}

type service struct {
	clientSet *kubernetes.Clientset
}

func NewService(client *kubernetes.Clientset) (Service, error) {
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

func (s service) Services(namespace string) Services {
	return s.Services(namespace)
}
