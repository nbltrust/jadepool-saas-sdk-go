package jadepoolsaas

const (
	defaultAddr = "http://saas.jadepool.io:8092"
)

type client interface {
	getKey() string
	getKeyHeaderName() string
	getSecret() string
	getAddr() string
}
