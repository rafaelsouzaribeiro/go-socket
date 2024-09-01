package serverslice

import (
	"fmt"
	"log"

	Connection "github.com/rafaelsouzaribeiro/go-socket/internal/infra/web"
	"github.com/rafaelsouzaribeiro/go-socket/pkg/factory"
	"github.com/rafaelsouzaribeiro/go-socket/pkg/global"
)

func RunServer() {
	connect := Connection.New("localhost", "8585")
	listener, err := connect.ConnectionServer()

	if err != nil {
		panic(err)
	}

	defer listener.Close()

	fmt.Printf("Server started at %s:%s \n", connect.Host, connect.Port)
	factories := factory.NewServer(factory.Slice)
	channel := make(chan global.Custom)
	factories.ChannelPerson = channel

	go func() {
		for p := range channel {
			fmt.Printf("Message received. Name: %s, Age: %d\n", p.Name, p.Age)
		}
	}()

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Fatal(err)
		}

		go factories.GetServer(connect, conn)

	}
}
