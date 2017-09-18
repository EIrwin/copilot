package resource

import (
	"encoding/json"

	"github.com/eirwin/copilot/pkg/format"
	"github.com/eirwin/copilot/pkg/internal/k8s"
)

type DeploymentRequest struct {
	namespace string
	service   k8s.Service
}

type Deployment struct {
	name      string `json:"name"`
	desired   string `json:"desired"`
	current   string `json:"current`
	upToDate  string `json:"upToDate"`
	available string `json:"available"`
	age       string `json:"age"`
}

type DeploymentResult struct {
	deployments []Deployment
}

func (r DeploymentRequest) Get(opts Options) Result {
	//r.clientSet.AppsV1beta1().Deployments(r.namespace).List()
	return DeploymentResult{
		deployments: []Deployment{},
	}
}

func (r DeploymentRequest) Logs(opts Options) Result {
	return DeploymentResult{
		deployments: []Deployment{},
	}
}

func (r DeploymentRequest) Status(opts Options) Result {
	return DeploymentResult{
		deployments: []Deployment{},
	}
}

func (r DeploymentResult) JSON() (string, error) {
	json, err := json.MarshalIndent(r.deployments, defaultJSONPrefix, defaultJSONIndent)
	if err != nil {
		return "", err
	}
	return string(json), nil
}

func (r DeploymentResult) Columns() (string, error) {
	formatter := format.NewColumnFormatter()
	headers := r.Headers()
	data := r.Data()
	cols := formatter.Format(headers, data)
	return cols, nil
}

func (r DeploymentResult) Headers() []string {
	return []string{
		"NAME",
		"DESIRED",
		"CURRENT",
		"UP-TO-DATE",
		"AVAILABLE",
		"AGE",
	}
}

func (r DeploymentResult) Data() [][]string {
	var data [][]string
	for _, d := range r.deployments {
		data = append(data, []string{
			d.name,
			d.desired,
			d.current,
			d.upToDate,
			d.available,
			d.age,
		})
	}
	return data
}
