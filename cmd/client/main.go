package main

import (
	Connection "github.com/rafaelsouzaribeiro/go-socket/internal/infra/web"
)

type Send struct {
	Message string
}

func main() {
	connect := Connection.New("localhost", "8585")
	conn, err := connect.ConnectionClient()

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	send := Send{Message: "Hello world 1"}

	messageBytes := []byte(send.Message)

	_, err = conn.Write(messageBytes)
	if err != nil {
		panic(err)
	}

}
