package k8s

import (
	"fmt"

	"strings"
	"time"

	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/api/v1"
)

type Pod struct {
	Name     string
	Ready    string
	Status   string
	Restarts string
	Age      string
}

type PodsGetter interface {
	Pods(namespace string) PodsInterface
}

type PodsInterface interface {
	List(opts ListOptions) ([]Pod, error)
}

type podsGetter struct {
	clientSet *kubernetes.Clientset
}

func (p podsGetter) Pods(namespace string) PodsInterface {
	return newPods(namespace, p.clientSet)
}

type pods struct {
	clientSet *kubernetes.Clientset
	namespace string
}

func newPods(namespace string, clientSet *kubernetes.Clientset) pods {
	return pods{
		clientSet: clientSet,
		namespace: namespace,
	}
}

func (p pods) List(opts ListOptions) ([]Pod, error) {
	var pods []Pod
	list, err := p.clientSet.CoreV1().Pods(p.namespace).List(meta_v1.ListOptions{
		LabelSelector: opts.LabelSelector,
		FieldSelector: opts.FieldSeletor,
	})
	if err != nil {
		return pods, err
	}

	for _, p := range list.Items {
		status := p.Status
		pod := Pod{
			Name:     p.Name,
			Ready:    calculateReady(status),
			Status:   string(status.Phase),
			Restarts: calculateRestarts(status),
			Age:      calculateAge(p.ObjectMeta),
		}

		pods = append(pods, pod)
	}
	return pods, nil
}

func calculateReady(status v1.PodStatus) string {
	var ready int32
	var total int32
	for _, s := range status.ContainerStatuses {
		total += 1
		if s.Ready {
			ready += 1
		}
	}
	return fmt.Sprintf("%v/%v", ready, total)
}

func calculateRestarts(status v1.PodStatus) string {
	var restarts int32
	for _, s := range status.ContainerStatuses {
		restarts += s.RestartCount
	}
	return fmt.Sprintf("%v", restarts)
}

func calculateAge(meta meta_v1.ObjectMeta) string {
	duration := time.Now().Sub(meta.CreationTimestamp.Time)
	return shortDur(duration)
}

func shortDur(d time.Duration) string {
	s := d.String()
	if strings.HasSuffix(s, "m0s") {
		s = s[:len(s)-2]
	}
	if strings.HasSuffix(s, "h0m") {
		s = s[:len(s)-2]
	}
	return s
}
