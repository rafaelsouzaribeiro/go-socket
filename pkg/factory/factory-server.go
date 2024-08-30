package factory

import Connection "github.com/rafaelsouzaribeiro/go-socket/internal/infra/web"

type FactoryServer struct {
	types         string
	ChannelPerson chan Connection.Person
	ChannelString chan string
}

const (
	Slice  = "slice"
	Struct = "struct"
	String = "string"
)

func NewServer(types string) *FactoryServer {
	return &FactoryServer{
		types: types,
	}
}
