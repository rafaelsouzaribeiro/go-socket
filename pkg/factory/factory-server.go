package factory

import "github.com/rafaelsouzaribeiro/go-socket/pkg/global"

type FactoryServer struct {
	types         string
	ChannelPerson chan global.Custom
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
