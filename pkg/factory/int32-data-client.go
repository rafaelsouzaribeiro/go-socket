package factory

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func (f *FactoryClient) getDataIntegerClient() ([]byte, error) {

	buf := new(bytes.Buffer)
	var intToWrite int32 = int32(*f.TypeInt)
	err := binary.Write(buf, binary.LittleEndian, intToWrite)
	if err != nil {
		return nil, fmt.Errorf("failed to convert integer to bytes: %v", err)
	}
	return buf.Bytes(), nil
}
