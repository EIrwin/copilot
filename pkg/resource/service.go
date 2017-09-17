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

func (s ServiceResult) JSON() (string, error) {
	json, err := jsoniter.MarshalIndent(s.services, defaultJSONPrefix, defaultJSONIndent)
	if err != nil {
		return "", err
	}
	return string(json), nil
}

func (s ServiceResult) Columns() (string, error) {
	formatter := format.NewColumnFormatter()
	headers := serviceColumnHeaders()
	data := serviceColumnData(s)
	cols := formatter.Format(headers, data)
	return cols, nil
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

func serviceColumnHeaders() []string {
	return []string{
		"name",
		"clusterIP",
		"externalIP",
		"ports",
		"age",
	}
}

func serviceColumnData(result ServiceResult) [][]string {
	var data [][]string
	for _, d := range result.services {
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
