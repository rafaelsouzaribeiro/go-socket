package main

import (
	Connection "github.com/rafaelsouzaribeiro/go-socket/internal/infra/web"
	"github.com/rafaelsouzaribeiro/go-socket/pkg/factory"
	"github.com/rafaelsouzaribeiro/go-socket/pkg/global"
)

func main() {
	connect := Connection.New("localhost", "8585")
	conn, err := connect.ConnectionClient()

	if err != nil {
		panic(err)
	}

	defer conn.Close()
	people := []global.Custom{
		{Name: "Rafael", Age: 38},
		{Name: "Maria", Age: 30},
		{Name: "João", Age: 25},
	}
	factories := factory.NewClient(global.Custom{}, people, "")
	buffer, err := factories.GetClient()

	if err != nil {
		panic(err)
	}

	_, err = conn.Write(buffer)
	if err != nil {
		panic(err)
	}

}
