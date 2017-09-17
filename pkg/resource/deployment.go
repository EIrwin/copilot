package resource

import (
	"github.com/eirwin/copilot/pkg/format"
	"github.com/eirwin/copilot/pkg/internal/k8s"
	"github.com/json-iterator/go"
)

type DeploymentRequest struct {
	namespace string
	service   k8s.Service
}

type Deployment struct {
	name      string `json:"name"`
	desired   int    `json:"desired"`
	current   int    `json:"current`
	upToDate  int    `json:"upToDate"`
	available int    `json:"available"`
	age       string `json:"age"`
}

type DeploymentResult struct {
	deployments []Deployment
}

func (d DeploymentResult) JSON() (string, error) {
	json, err := jsoniter.MarshalIndent(d.deployments, defaultJSONPrefix, defaultJSONIndent)
	if err != nil {
		return "", err
	}
	return string(json), nil
}

func (d DeploymentResult) Columns() (string, error) {
	formatter := format.NewColumnFormatter()
	headers := deploymentColumnHeaders()
	data := deploymentColumnData(d)
	cols := formatter.Format(headers, data)
	return cols, nil
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

func deploymentColumnHeaders() []string {
	return []string{
		"name",
		"desired",
		"current",
		"upToDate",
		"available",
		"age",
	}
}

func deploymentColumnData(result DeploymentResult) [][]string {
	var data [][]string
	for _, d := range result.deployments {
		data = append(data, []string{
			d.name,
			string(d.desired),
			string(d.current),
			string(d.upToDate),
			string(d.available),
			string(d.age),
		})
	}
	return data
}
