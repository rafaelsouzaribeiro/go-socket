package factory

import (
	"github.com/rafaelsouzaribeiro/go-socket/pkg/global"
)

type FactoryClient struct {
	typeStruct global.Custom
	typeSlice  []global.Custom
	typeString string
}

func isPersonEmpty(p global.Custom) bool {
	return p == global.Custom{}
}

func NewClient(typeStruct global.Custom, typeSlice []global.Custom, typeString string) *FactoryClient {
	return &FactoryClient{
		typeStruct: typeStruct,
		typeSlice:  typeSlice,
		typeString: typeString,
	}
}
