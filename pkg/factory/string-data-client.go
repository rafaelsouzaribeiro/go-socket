package factory

func (f *FactoryClient) getDataStringClient() ([]byte, error) {

	return []byte(f.typeString), nil
}
