package resource

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

func NewRequest(name string) Request {

	var req Request
	switch name {
	case "pod":
		req = PodRequest{}
		break
	case "deployment":
		req = DeploymentRequest{}
		break
	case "service":
		req = ServiceRequest{}
		break
	}

	return req
}
