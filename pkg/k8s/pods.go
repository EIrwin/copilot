package k8s

type PodsGetter interface {
	Pods(namespace string) Pods
}

type Pods interface {
	List()
}
