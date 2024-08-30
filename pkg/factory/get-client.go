package factory

import (
	"fmt"
)

func (f *FactoryClient) GetClient() ([]byte, error) {
	if !isPersonEmpty(f.typeStruct) {
		return f.getDataStructClient()
	}

	if f.typeSlice != nil {
		return f.getDataSliceClient()
	}

	if f.typeString != "" {
		return f.getDataStringClient()
	}

	return nil, fmt.Errorf("invalid function")
}
