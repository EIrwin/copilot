package resource

import "github.com/eirwin/copilot/pkg/internal/k8s"

type PodRequest struct {
	namespace string
	service   k8s.Service
}

type Pod struct {
	name     string `json:"name"`
	status   string `json:"status`
	restarts int    `json:"restarts`
	age      string `json:"age"`
}

type PodResult struct {
	pods []Pod
}

func (r PodRequest) Get(opts Options) GetResult {
	return GetResult{}
}

func (r PodRequest) Logs(opts Options) LogsResult {
	return LogsResult{}
}

func (r PodRequest) Status(opts Options) StatusResult {
	return StatusResult{}
}
