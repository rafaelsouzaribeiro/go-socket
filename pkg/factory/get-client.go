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

	return nil, fmt.Errorf("invalid function")
}
