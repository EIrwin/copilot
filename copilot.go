package copilot

import "github.com/eirwin/copilot/pkg/resource"

type Service interface {
	Run(cmd Command) (string, error)
}

type service struct {
}

type Config struct {
}

func NewService(config Config) (Service, error) {
	return service{}, nil
}

func (s service) Run(cmd Command) (string, error) {

	switch cmd.action {
	case "get":
		req := resource.NewRequest(cmd.resource).Get(nil)
		break
	case "logs":
		req := resource.NewRequest(cmd.resource).Logs(nil)
		break
	case "status":
		req := resource.NewRequest(cmd.resource).Status(nil)
		break
	}

	//TODO: format output

	return nil
}
