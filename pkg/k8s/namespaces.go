package k8s

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Namespace struct {
	Name   string
	Status string
	Age    string
}

type NamespacesGetter interface {
	Namespaces() NamespacesInterface
}

type NamespacesInterface interface {
	List(opts ListOptions) ([]Namespace, error)
}

type namespacesGetter struct {
	clientSet *kubernetes.Clientset
}

func (n namespacesGetter) Namespaces() NamespacesInterface {
	return newNamespaces(n.clientSet)
}

type namespaces struct {
	clientSet *kubernetes.Clientset
}

func newNamespaces(clientSet *kubernetes.Clientset) namespaces {
	return namespaces{
		clientSet: clientSet,
	}
}

func (n namespaces) List(opts ListOptions) ([]Namespace, error) {
	var namespaces []Namespace
	list, err := n.clientSet.Namespaces().List(meta_v1.ListOptions{
		LabelSelector: opts.LabelSelector,
		FieldSelector: opts.FieldSeletor,
	})
	if err != nil {
		return namespaces, err
	}

	for _, n := range list.Items {
		namespace := Namespace{
			Name:   n.Name,
			Status: string(n.Status.Phase),
			Age:    calculateAge(n.ObjectMeta),
		}

		namespaces = append(namespaces, namespace)
	}

	return namespaces, nil
}
