package factory

import Connection "github.com/rafaelsouzaribeiro/go-socket/internal/infra/web"

type FactoryServer struct {
	typeStruct bool
	typeSlice  bool
	Channel    chan Connection.Person
}

func NewServer(typeStruct bool, typeSlice bool) *FactoryServer {
	return &FactoryServer{
		typeStruct: typeStruct,
		typeSlice:  typeSlice,
	}
}
