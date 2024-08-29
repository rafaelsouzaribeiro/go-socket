package factory

import Connection "github.com/rafaelsouzaribeiro/go-socket/internal/infra/web"

type FactoryClient struct {
	typeStruct Connection.Person
	typeSlice  []Connection.Person
}

func isPersonEmpty(p Connection.Person) bool {
	return p == Connection.Person{}
}

func NewClient(typeStruct Connection.Person, typeSlice []Connection.Person) *FactoryClient {
	return &FactoryClient{
		typeStruct: typeStruct,
		typeSlice:  typeSlice,
	}
}
