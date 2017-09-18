package resource

import (
	"github.com/eirwin/copilot/pkg/format"
	"github.com/json-iterator/go"
)

type ServiceRequest struct{}

type Service struct {
	name       string `json:"name"`
	clusterIP  string `json:"clusterIp"`
	externalIP string `json:"externalIp"`
	ports      string `json:"ports"`
	age        string `json:"age"`
}

type ServiceResult struct {
	services []Service
}

func (r ServiceRequest) Get(opts Options) Result {
	return ServiceResult{
		services: []Service{},
	}
}

func (r ServiceRequest) Logs(opts Options) Result {
	return ServiceResult{
		services: []Service{},
	}
}

func (r ServiceRequest) Status(opts Options) Result {
	return ServiceResult{
		services: []Service{},
	}
}

func (r ServiceResult) JSON() (string, error) {
	json, err := jsoniter.MarshalIndent(r.services, defaultJSONPrefix, defaultJSONIndent)
	if err != nil {
		return "", err
	}
	return string(json), nil
}

func (r ServiceResult) Columns() (string, error) {
	formatter := format.NewColumnFormatter()
	headers := r.Headers()
	data := r.Data()
	cols := formatter.Format(headers, data)
	return cols, nil
}

func (r ServiceResult) Headers() []string {
	return []string{
		"NAME",
		"CLUSTER-IP",
		"EXTERNAL-IP",
		"PORTS",
		"AGE",
	}
}

func (r ServiceResult) Data() [][]string {
	var data [][]string
	for _, d := range r.services {
		data = append(data, []string{
			d.name,
			d.clusterIP,
			d.externalIP,
			d.ports,
			d.age,
		})
	}
	return data
}
