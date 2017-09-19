package k8s

import (
	"fmt"

	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/apis/apps/v1beta1"
)

type Deployment struct {
	Name      string
	Desired   string
	Current   string
	UpToDate  string
	Available string
	Age       string
}

type DeploymentsGetter interface {
	Deployments(namespace string) DeploymentsInterface
}

type DeploymentsInterface interface {
	List(opts ListOptions) ([]Deployment, error)
}

type deploymentsGetter struct {
	clientSet *kubernetes.Clientset
}

func (d deploymentsGetter) Deployments(namespace string) DeploymentsInterface {
	return newDeployments(namespace, d.clientSet)
}

type deployments struct {
	clientSet *kubernetes.Clientset
	namespace string
}

func newDeployments(namespace string, clientSet *kubernetes.Clientset) deployments {
	return deployments{
		clientSet: clientSet,
		namespace: namespace,
	}
}

func (d deployments) List(opts ListOptions) ([]Deployment, error) {
	var deployments []Deployment
	list, err := d.clientSet.AppsV1beta1().Deployments(d.namespace).List(meta_v1.ListOptions{
		LabelSelector: opts.LabelSelector,
		FieldSelector: opts.FieldSeletor,
	})
	if err != nil {
		return deployments, err
	}
	for _, d := range list.Items {
		status := d.Status
		deployment := Deployment{
			Name:      d.Name,
			Desired:   calculateDesired(status),
			Current:   calculateCurrent(status),
			UpToDate:  calculateUpToDate(status),
			Available: calculateAvailable(status),
			Age:       calculateAge(d.ObjectMeta),
		}

		deployments = append(deployments, deployment)
	}

	return deployments, nil
}

func calculateDesired(status v1beta1.DeploymentStatus) string {
	return fmt.Sprintf("%v", status.Replicas)
}

func calculateCurrent(status v1beta1.DeploymentStatus) string {
	return fmt.Sprintf("%v", status.ReadyReplicas)
}

func calculateAvailable(status v1beta1.DeploymentStatus) string {
	return fmt.Sprintf("%v", status.AvailableReplicas)
}

func calculateUpToDate(status v1beta1.DeploymentStatus) string {
	return fmt.Sprintf("%v", status.UpdatedReplicas)
}
