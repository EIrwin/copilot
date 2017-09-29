package k8s

import (
	"fmt"

	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/apis/extensions/v1beta1"
)

type ReplicaSet struct {
	Name    string
	Desired string
	Current string
	Ready   string
	Age     string
}

type ReplicaSetsGetter interface {
	ReplicaSets(namespace string) ReplicaSetsInterface
}

type ReplicaSetsInterface interface {
	List(opts ListOptions) ([]ReplicaSet, error)
}

type replicaSetsGetter struct {
	clientSet *kubernetes.Clientset
}

func (r replicaSetsGetter) ReplicaSets(namespace string) ReplicaSetsInterface {
	return newReplicaSets(namespace, r.clientSet)
}

type replicaSets struct {
	clientSet *kubernetes.Clientset
	namespace string
}

func newReplicaSets(namespace string, clientSet *kubernetes.Clientset) replicaSets {
	return replicaSets{
		clientSet: clientSet,
		namespace: namespace,
	}
}

func (r replicaSets) List(opts ListOptions) ([]ReplicaSet, error) {
	var replicaSets []ReplicaSet
	list, err := r.clientSet.ReplicaSets(r.namespace).List(meta_v1.ListOptions{
		LabelSelector: opts.LabelSelector,
		FieldSelector: opts.FieldSeletor,
	})
	if err != nil {
		return replicaSets, err

	}
	for _, rs := range list.Items {
		status := rs.Status
		replicaSet := ReplicaSet{
			Name:    rs.Name,
			Desired: calculateDesiredReplicaSets(status),
			Current: calculateCurrentReplicaSets(status),
			Ready:   calculateReadyReplicaSets(status),
			Age:     calculateAge(rs.ObjectMeta),
		}

		replicaSets = append(replicaSets, replicaSet)
	}

	return replicaSets, nil
}

func calculateDesiredReplicaSets(status v1beta1.ReplicaSetStatus) string {
	return fmt.Sprintf("%v", status.Replicas)
}

func calculateCurrentReplicaSets(status v1beta1.ReplicaSetStatus) string {
	return fmt.Sprintf("%v", status.AvailableReplicas)
}

func calculateReadyReplicaSets(status v1beta1.ReplicaSetStatus) string {
	return fmt.Sprintf("%v", status.ReadyReplicas)
}
