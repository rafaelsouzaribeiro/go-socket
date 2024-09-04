package factory

import (
	"fmt"
)

func (f *FactoryClient) GetClient() ([]byte, error) {
	if !isPersonEmpty(f.TypeStruct) {
		return f.getDataStructClient()
	}

	if f.TypeSlice != nil {
		return f.getDataSliceClient()
	}

	if f.TypeString != "" {
		return f.getDataStringClient()
	}

	if f.TypeInt != nil {
		return f.getDataIntegerClient()

	}

	return nil, fmt.Errorf("invalid function")
}
