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

	if cmd.Resource == "" {
		return "", NewErrInvalidRunArgument("resource")
	}

	if cmd.Namespace == "" {
		return "", NewErrInvalidRunArgument("namespace")
	}

	factory := resource.NewRequestFactory(s.kubernetes)

	var result resource.Result
	switch cmd.Operation {
	case "get":
		result = factory.NewRequest(cmd.Resource, cmd.Namespace).Get(nil)
		break
	case "logs":
		result = factory.NewRequest(cmd.Resource, cmd.Namespace).Logs(nil)
		break
	case "status":
		result = factory.NewRequest(cmd.Resource, cmd.Namespace).Status(nil)
		break
	}

	output, err := formatOutput(result, cmd.Output)
	if err != nil {
		return "", err
	}

	return output, nil
}

func formatOutput(result resource.Result, format string) (string, error) {
	var output string
	switch format {
	case "json":
		return result.JSON()
	case "columns":
		return result.Columns()
	default:
		return result.Columns()
	}
	return output, nil
}
