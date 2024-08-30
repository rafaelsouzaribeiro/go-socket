package factory

import Connection "github.com/rafaelsouzaribeiro/go-socket/internal/infra/web"

type FactoryClient struct {
	typeStruct Connection.Person
	typeSlice  []Connection.Person
	typeString string
}

func isPersonEmpty(p Connection.Person) bool {
	return p == Connection.Person{}
}

func NewClient(typeStruct Connection.Person, typeSlice []Connection.Person, typeString string) *FactoryClient {
	return &FactoryClient{
		typeStruct: typeStruct,
		typeSlice:  typeSlice,
		typeString: typeString,
	}
}
