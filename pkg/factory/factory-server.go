package factory

import "github.com/rafaelsouzaribeiro/go-socket/pkg/global"

type FactoryServer struct {
	types         string
	ChannelCustom chan global.Custom
	ChannelString chan string
	ChannelInt    chan int32
}

const (
	Slice  = "slice"
	Struct = "struct"
	String = "string"
	Int32  = "Int32"
)

func NewServer(types string) *FactoryServer {
	return &FactoryServer{
		types: types,
	}
}
