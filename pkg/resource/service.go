package resource

import (
	"github.com/eirwin/copilot/pkg/format"
	"github.com/eirwin/copilot/pkg/k8s"
	jsoniter "github.com/json-iterator/go"
)

type ServiceRequest struct {
	namespace string
	service   k8s.Kubernetes
}

type Service struct {
	Name       string `json:"name"`
	ClusterIP  string `json:"clusterIp"`
	ExternalIP string `json:"externalIp"`
	Ports      string `json:"ports"`
	Age        string `json:"age"`
}

type ServiceResult struct {
	Services []Service `json:"services"`
}

func (r ServiceRequest) Get(opts Options) Result {
	var services []Service
	serviceList, _ := r.service.Services(r.namespace).List(k8s.ListOptions{})

	for _, s := range serviceList {
		service := Service{
			Name:       s.Name,
			ClusterIP:  s.ClusterIP,
			ExternalIP: s.ExternalIP,
			Ports:      s.Ports,
			Age:        s.Age,
		}

		services = append(services, service)
	}

	return ServiceResult{
		Services: services,
	}
}

func (r ServiceRequest) Logs(opts Options) Result {
	return ServiceResult{
		Services: []Service{},
	}
}

func (r ServiceRequest) Status(opts Options) Result {
	return ServiceResult{
		Services: []Service{},
	}
}

func (r ServiceResult) JSON() (string, error) {
	json, err := jsoniter.MarshalIndent(r, defaultJSONPrefix, defaultJSONIndent)
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
	for _, d := range r.Services {
		data = append(data, []string{
			d.Name,
			d.ClusterIP,
			d.ExternalIP,
			d.Ports,
			d.Age,
		})
	}
	return data
}
