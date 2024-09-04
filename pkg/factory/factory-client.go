package factory

import (
	"github.com/rafaelsouzaribeiro/go-socket/pkg/global"
)

type FactoryClient struct {
	TypeStruct global.Custom
	TypeSlice  []global.Custom
	TypeString string
	TypeInt    *int
}

func isPersonEmpty(p global.Custom) bool {
	return p == global.Custom{}
}

func NewClient(factory FactoryClient) *FactoryClient {
	return &factory
}
