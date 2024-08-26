package main

import (
	"bytes"
	"encoding/gob"

	Connection "github.com/rafaelsouzaribeiro/go-socket/internal/infra/web"
)

func main() {
	connect := Connection.New("localhost", "8585")
	conn, err := connect.ConnectionClient()

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	send := Connection.Person{Name: "Rafael", Age: 38}

	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err = encoder.Encode(send)

	if err != nil {
		panic(err)
	}

	_, err = conn.Write(buffer.Bytes())
	if err != nil {
		panic(err)
	}

}
