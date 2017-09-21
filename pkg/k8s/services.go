package k8s

import (
	"fmt"

	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/api/v1"
)

type Service struct {
	Name       string
	ClusterIP  string
	ExternalIP string
	Ports      string
	Age        string
}
type ServicesGetter interface {
	Services(namespace string) ServicesInterface
}

type ServicesInterface interface {
	List(opts ListOptions) ([]Service, error)
}

type servicesGetter struct {
	clientSet *kubernetes.Clientset
}

func (s servicesGetter) Services(namespace string) ServicesInterface {
	return newServices(namespace, s.clientSet)
}

type services struct {
	clientSet *kubernetes.Clientset
	namespace string
}

func newServices(namespace string, clientSet *kubernetes.Clientset) services {
	return services{
		clientSet: clientSet,
		namespace: namespace,
	}
}

func (s services) List(opts ListOptions) ([]Service, error) {
	var services []Service
	list, err := s.clientSet.Services(s.namespace).List(meta_v1.ListOptions{
		LabelSelector: opts.LabelSelector,
		FieldSelector: opts.FieldSeletor,
	})
	if err != nil {
		return services, err
	}

	for _, s := range list.Items {
		service := Service{
			Name:       s.Name,
			ClusterIP:  s.Spec.ClusterIP,
			ExternalIP: formatExternalIPs(s.Spec.ExternalIPs),
			Ports:      formatPorts(s.Spec.Ports),
			Age:        calculateAge(s.ObjectMeta),
		}

		services = append(services, service)
	}
	return services, nil
}

func formatExternalIPs(ips []string) string {
	//TODO need to determine extended format
	var externalIPs string
	for _, ip := range ips {
		if externalIPs != "" {
			externalIPs += ","
		}
		externalIPs += ip
	}

	if externalIPs == "" {
		externalIPs = "<none>"
	}

	return externalIPs
}

func formatPorts(servicePorts []v1.ServicePort) string {
	//TODO need to determine extended format
	var ports string
	for _, p := range servicePorts {
		if ports != "" {
			ports += ","
		}

		ports += fmt.Sprintf("%v/%v", p.Port, p.Protocol)
	}
	return ports
}
