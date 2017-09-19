package resource

import (
	"encoding/json"

	"github.com/eirwin/copilot/pkg/format"
	"github.com/eirwin/copilot/pkg/k8s"
)

type DeploymentRequest struct {
	namespace string
	service   k8s.Kubernetes
}

type Deployment struct {
	Name      string `json:"name"`
	Desired   string `json:"desired"`
	Current   string `json:"current`
	UpToDate  string `json:"upToDate"`
	Available string `json:"available"`
	Age       string `json:"age"`
}

type DeploymentResult struct {
	Deployments []Deployment
}

func (r DeploymentRequest) Get(opts Options) Result {
	var deployments []Deployment
	deploymentList, _ := r.service.Deployments(r.namespace).List(k8s.ListOptions{})

	for _, d := range deploymentList {
		deployment := Deployment{
			Name:      d.Name,
			Desired:   d.Desired,
			Current:   d.Current,
			UpToDate:  d.UpToDate,
			Available: d.Available,
			Age:       d.Age,
		}
		deployments = append(deployments, deployment)
	}

	return DeploymentResult{
		Deployments: deployments,
	}
}

func (r DeploymentRequest) Logs(opts Options) Result {
	return DeploymentResult{
		Deployments: []Deployment{},
	}
}

func (r DeploymentRequest) Status(opts Options) Result {
	return DeploymentResult{
		Deployments: []Deployment{},
	}
}

func (r DeploymentResult) JSON() (string, error) {
	json, err := json.MarshalIndent(r.Deployments, defaultJSONPrefix, defaultJSONIndent)
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
	for _, d := range r.Deployments {
		data = append(data, []string{
			d.Name,
			d.Desired,
			d.Current,
			d.UpToDate,
			d.Available,
			d.Age,
		})
	}
	return data
}
