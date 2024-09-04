package int32

import (
	"sync"
	"testing"
	"time"

	Connection "github.com/rafaelsouzaribeiro/go-socket/internal/infra/web"
	"github.com/rafaelsouzaribeiro/go-socket/pkg/factory"
	serverint32 "github.com/rafaelsouzaribeiro/go-socket/test/server-int32"
)

var Once sync.Once

func BenchmarkClient(b *testing.B) {
	Once.Do(func() {
		go serverint32.RunServer()
		time.Sleep(1 * time.Second)
	})

	for i := 0; i < b.N; i++ {
		connect := Connection.New("localhost", "8585")
		conn, err := connect.ConnectionClient()

		if err != nil {
			panic(err)
		}

		defer conn.Close()
		valueint := 5

		factories := factory.NewClient(factory.FactoryClient{
			TypeInt: &valueint,
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
