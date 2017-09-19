package k8s

import "k8s.io/client-go/kubernetes"

type Kubernetes interface {
	DeploymentsGetter
	PodsGetter
	ServicesGetter
}

type service struct {
	clientSet *kubernetes.Clientset
}

func NewService(client *kubernetes.Clientset) (Kubernetes, error) {
	return service{
		clientSet: client,
	}, nil
}

func (s service) Deployments(namespace string) DeploymentsInterface {
	return deployments{
		namespace: namespace,
		clientSet: s.clientSet,
	}
}

func (s service) Pods(namespace string) PodsInterface {
	return pods{
		namespace: namespace,
		clientSet: s.clientSet,
	}
}

func (s service) Services(namespace string) ServicesInterface {
	return services{
		namespace: namespace,
		clientSet: s.clientSet,
	}
}
