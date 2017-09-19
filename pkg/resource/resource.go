package resource

import "github.com/eirwin/copilot/pkg/k8s"

type Getter interface {
	Get(opt Options) Result
}

type Logger interface {
	Logs(opt Options) Result
}

type Status interface {
	Status(opt Options) Result
}

type Request interface {
	Getter
	Logger
	Status
}

type Options map[string]string

type RequestFactory struct {
	service k8s.Kubernetes
}

func NewRequestFactory(kubernetes k8s.Kubernetes) RequestFactory {
	return RequestFactory{
		service: kubernetes,
	}
}

func (b RequestFactory) NewRequest(name, namespace string) Request {

	var req Request
	switch name {
	case "pod":
		req = PodRequest{
			service:   b.service,
			namespace: namespace,
		}
		break
	case "deployment":
		req = DeploymentRequest{
			service:   b.service,
			namespace: namespace,
		}
		break
	case "service":
		req = ServiceRequest{
			service:   b.service,
			namespace: namespace,
		}
		break
	}

	return req
}
