package serverfolder

import (
	"fmt"
	"log"

	Connection "github.com/rafaelsouzaribeiro/go-socket/internal/infra/web"
	"github.com/rafaelsouzaribeiro/go-socket/pkg/factory"
)

func RunServer() {
	connect := Connection.New("localhost", "8585")
	listener, err := connect.ConnectionServer()

	if err != nil {
		panic(err)
	}

	defer listener.Close()

	fmt.Printf("Server started at %s:%s \n", connect.Host, connect.Port)
	factories := factory.NewServer(factory.Folder)

	channel := make(chan string)
	factories.ChannelString = channel

	go func() {
		for p := range channel {
			fmt.Printf("Message received: %s", p)
		}
	}()

	i := 0
	for {
		factories.OutputFolderFileName = fmt.Sprintf("../../cmd/received%d.zip", i)

		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Fatal(err)
		}

		go factories.GetServer(connect, conn)
		i++
	}

}
