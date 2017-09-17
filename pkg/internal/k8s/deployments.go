package k8s

type DeploymentsGetter interface {
	Deployments(namespace string) Deployments
}

type Deployments interface {
	List()
}
