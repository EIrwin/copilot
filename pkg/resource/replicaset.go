package resource

import (
	"github.com/eirwin/copilot/pkg/format"
	"github.com/eirwin/copilot/pkg/k8s"
	"github.com/json-iterator/go"
)

type ReplicaSetRequest struct {
	namespace string
	service   k8s.Kubernetes
}

type ReplicaSet struct {
	Name    string `json:"name"`
	Desired string `json:"desired"`
	Current string `json:"current"`
	Ready   string `json:"ready"`
	Age     string `json:"age"`
}

type ReplicaSetResult struct {
	ReplicaSets []ReplicaSet
}

func (r ReplicaSetRequest) Get(opts Options) Result {
	var replicaSets []ReplicaSet
	replicaSetList, _ := r.service.ReplicaSets(r.namespace).List(k8s.ListOptions{})

	for _, rs := range replicaSetList {
		replicaSet := ReplicaSet{
			Name:    rs.Name,
			Desired: rs.Desired,
			Current: rs.Current,
			Ready:   rs.Ready,
			Age:     rs.Age,
		}

		replicaSets = append(replicaSets, replicaSet)
	}

	return ReplicaSetResult{
		ReplicaSets: replicaSets,
	}
}

func (r ReplicaSetRequest) Logs(opts Options) Result {
	return ReplicaSetResult{
		ReplicaSets: []ReplicaSet{},
	}
}

func (r ReplicaSetRequest) Status(opts Options) Result {
	return ReplicaSetResult{
		ReplicaSets: []ReplicaSet{},
	}
}

func (r ReplicaSetResult) JSON() (string, error) {
	json, err := jsoniter.MarshalIndent(r.ReplicaSets, defaultJSONPrefix, defaultJSONIndent)
	if err != nil {
		return "", err
	}
	return string(json), nil
}

func (r ReplicaSetResult) Columns() (string, error) {
	formatter := format.NewColumnFormatter()
	headers := r.Headers()
	data := r.Data()
	cols := formatter.Format(headers, data)
	return cols, nil
}

func (r ReplicaSetResult) Headers() []string {
	return []string{
		"NAME",
		"DESIRED",
		"CURRENT",
		"READY",
		"AGE",
	}
}

func (r ReplicaSetResult) Data() [][]string {
	var data [][]string
	for _, d := range r.ReplicaSets {
		data = append(data, []string{
			d.Name,
			d.Desired,
			d.Current,
			d.Ready,
			d.Age,
		})
	}
	return data
}
