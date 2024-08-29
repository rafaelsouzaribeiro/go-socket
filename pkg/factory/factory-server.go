package factory

type FactoryServer struct {
	typeStruct bool
	typeSlice  bool
}

func NewServer(typeStruct bool, typeSlice bool) *FactoryServer {
	return &FactoryServer{
		typeStruct: typeStruct,
		typeSlice:  typeSlice,
	}
}
