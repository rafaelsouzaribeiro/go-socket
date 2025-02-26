package structs

import (
	"sync"
	"testing"
	"time"

	Connection "github.com/rafaelsouzaribeiro/go-socket/internal/infra/web"
	"github.com/rafaelsouzaribeiro/go-socket/pkg/factory"
	"github.com/rafaelsouzaribeiro/go-socket/pkg/global"
	serverstruct "github.com/rafaelsouzaribeiro/go-socket/test/server-struct"
)

var Once sync.Once

func BenchmarkClient(b *testing.B) {
	Once.Do(func() {
		go serverstruct.RunServer()
		time.Sleep(1 * time.Second)
	})

	for i := 0; i < b.N; i++ {
		connect := Connection.New("localhost", "8585")
		conn, err := connect.ConnectionClient()

		if err != nil {
			panic(err)
		}

		defer conn.Close()

		factories := factory.NewClient(factory.FactoryClient{
			TypeStruct: global.Custom{Name: "Rafael", Age: 38},
		})

		buffer, err := factories.GetClient()

		if err != nil {
			panic(err)
		}

		_, err = conn.Write(buffer)
		if err != nil {
			panic(err)
		}
	}
}
