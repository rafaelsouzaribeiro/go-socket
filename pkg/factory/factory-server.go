package factory

import Connection "github.com/rafaelsouzaribeiro/go-socket/internal/infra/web"

type FactoryServer struct {
	types   string
	Channel chan Connection.Person
}

const (
	Slice  = "slice"
	Struct = "struct"
)

func NewServer(types string) *FactoryServer {
	return &FactoryServer{
		types: types,
	}
}
