package resource

import (
	"encoding/json"

	"github.com/eirwin/copilot/pkg/format"
	"github.com/eirwin/copilot/pkg/k8s"
)

type NamespaceRequest struct {
	service k8s.Kubernetes
}

type Namespace struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Age    string `json:"age"`
}

type NamespaceResult struct {
	Namespaces []Namespace
}

func (r NamespaceRequest) Get(opts Options) Result {
	var namespaces []Namespace
	namespacesList, _ := r.service.Namespaces().List(k8s.ListOptions{})

	for _, n := range namespacesList {
		namespace := Namespace{
			Name:   n.Name,
			Status: n.Status,
			Age:    n.Age,
		}

		namespaces = append(namespaces, namespace)
	}

	return NamespaceResult{
		Namespaces: namespaces,
	}
}

func (r NamespaceRequest) Logs(opts Options) Result {
	return NamespaceResult{
		Namespaces: []Namespace{},
	}
}

func (r NamespaceRequest) Status(opts Options) Result {
	return NamespaceResult{
		Namespaces: []Namespace{},
	}
}

func (r NamespaceResult) JSON() (string, error) {
	json, err := json.MarshalIndent(r.Namespaces, defaultJSONPrefix, defaultJSONIndent)
	if err != nil {
		return "", err
	}
	return string(json), nil
}

func (r NamespaceResult) Columns() (string, error) {
	formatter := format.NewColumnFormatter()
	headers := r.Headers()
	data := r.Data()
	cols := formatter.Format(headers, data)
	return cols, nil
}

func (r NamespaceResult) Headers() []string {
	return []string{
		"NAME",
		"STATUS",
		"AGE",
	}
}

func (r NamespaceResult) Data() [][]string {
	var data [][]string
	for _, d := range r.Namespaces {
		data = append(data, []string{
			d.Name,
			d.Status,
			d.Age,
		})
	}
	return data
}
