package resource

import (
	"encoding/json"

	"github.com/eirwin/copilot/pkg/format"
	"github.com/eirwin/copilot/pkg/k8s"
)

type PodRequest struct {
	namespace string
	service   k8s.Kubernetes
}

type Pod struct {
	Name     string `json:"name"`
	Ready    string `json:"ready"`
	Status   string `json:"status`
	Restarts string `json:"restarts`
	Age      string `json:"age"`
}

type PodResult struct {
	Pods []Pod `json:"pods"`
}

func (r PodRequest) Get(opts Options) Result {
	var pods []Pod
	podList, _ := r.service.Pods(r.namespace).List(k8s.ListOptions{})

	for _, p := range podList {
		pod := Pod{
			Name:     p.Name,
			Ready:    p.Ready,
			Status:   p.Status,
			Restarts: p.Restarts,
			Age:      p.Age,
		}
		pods = append(pods, pod)
	}

	return PodResult{
		Pods: pods,
	}
}

func (r PodRequest) Logs(opts Options) Result {
	r.service.Pods(r.namespace)
	return PodResult{
		Pods: []Pod{},
	}
}

func (r PodRequest) Status(opts Options) Result {
	return PodResult{
		Pods: []Pod{},
	}
}

func (r PodResult) JSON() (string, error) {
	json, err := json.MarshalIndent(r, defaultJSONPrefix, defaultJSONIndent)
	if err != nil {
		return "", err
	}

	return string(json), nil
}

func (r PodResult) Columns() (string, error) {
	formatter := format.NewColumnFormatter()
	headers := r.Headers()
	data := r.Data()
	cols := formatter.Format(headers, data)
	return cols, nil
}

func (r PodResult) Headers() []string {
	return []string{
		"NAME",
		"READY",
		"STATUS",
		"RESTARTS",
		"AGE",
	}
}

func (r PodResult) Data() [][]string {
	var data [][]string
	for _, d := range r.Pods {
		data = append(data, []string{
			d.Name,
			d.Ready,
			d.Status,
			d.Restarts,
			d.Age,
		})
	}
	return data
}
