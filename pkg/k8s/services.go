package k8s

type ServicesGetter interface {
	Services(namespace string) Services
}

type Services interface {
	List()
}
