package copilot

import (
	"github.com/eirwin/copilot/pkg/k8s"
	"github.com/eirwin/copilot/pkg/resource"
)

type Service interface {
	Run(cmd Command) (string, error)
}

type service struct {
	kubernetes k8s.Kubernetes
}

func NewService(kubernetes k8s.Kubernetes) Service {
	return service{
		kubernetes: kubernetes,
	}
}

func (s service) Run(cmd Command) (string, error) {

	factory := resource.NewRequestFactory(s.kubernetes)

	var result resource.Result
	switch cmd.operation {
	case "get":
		result = factory.NewRequest(cmd.resource, cmd.namespace).Get(nil)
		break
	case "logs":
		result = factory.NewRequest(cmd.resource, cmd.namespace).Logs(nil)
		break
	case "status":
		result = factory.NewRequest(cmd.resource, cmd.namespace).Status(nil)
		break
	}

	cols, err := result.Columns()
	if err != nil {
		return "", err
	}

	return cols, nil
}
