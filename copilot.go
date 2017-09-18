package copilot

import (
	"log"

	"github.com/eirwin/copilot/pkg/k8s"
	"github.com/eirwin/copilot/pkg/resource"
)

type Service interface {
	Run(cmd Command) (string, error)
}

type service struct {
	kubernetes k8s.Service
}

func NewService(kubernetes k8s.Service) Service {
	return service{
		kubernetes: kubernetes,
	}
}

func (s service) Run(cmd Command) (string, error) {

	var result resource.Result
	switch cmd.operation {
	case "get":
		result = resource.NewRequest(cmd.resource).Get(nil)
		break
	case "logs":
		result = resource.NewRequest(cmd.resource).Logs(nil)
		break
	case "status":
		result = resource.NewRequest(cmd.resource).Status(nil)
		break
	}

	log.Println(result)

	//TODO: format output

	return "", nil
}
