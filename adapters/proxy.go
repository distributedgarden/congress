package adapters

type ProxyRepository interface {
	Create() (bool, error)
}
