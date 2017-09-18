package resource

import (
	"encoding/json"

	"github.com/eirwin/copilot/pkg/format"
	"github.com/eirwin/copilot/pkg/internal/k8s"
)

type PodRequest struct {
	namespace string
	service   k8s.Service
}

type Pod struct {
	name     string `json:"name"`
	status   string `json:"status`
	restarts string `json:"restarts`
	age      string `json:"age"`
}

type PodResult struct {
	pods []Pod
}

func (r PodRequest) Get(opts Options) Result {
	return PodResult{
		pods: []Pod{},
	}
}

func (r PodRequest) Logs(opts Options) Result {
	return PodResult{
		pods: []Pod{},
	}
}

func (r PodRequest) Status(opts Options) Result {
	return PodResult{
		pods: []Pod{},
	}
}

func (r PodResult) JSON() (string, error) {
	json, err := json.MarshalIndent(r.pods, defaultJSONPrefix, defaultJSONIndent)
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
		"DESIRED",
		"CURRENT",
		"UP-TO-DATE",
		"AVAILABLE",
		"AGE",
	}
}

func (r PodResult) Data() [][]string {
	var data [][]string
	for _, d := range r.pods {
		data = append(data, []string{
			d.name,
			d.status,
			d.restarts,
			d.age,
		})
	}
	return data
}
