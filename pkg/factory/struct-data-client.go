package factory

import (
	"bytes"
	"encoding/gob"
)

func (f *FactoryClient) getDataStructClient() ([]byte, error) {

	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(f.typeStruct)

	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
