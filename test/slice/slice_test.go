package slice

import (
	"sync"
	"testing"
	"time"

	Connection "github.com/rafaelsouzaribeiro/go-socket/internal/infra/web"
	"github.com/rafaelsouzaribeiro/go-socket/pkg/factory"
	"github.com/rafaelsouzaribeiro/go-socket/pkg/global"
	serverslice "github.com/rafaelsouzaribeiro/go-socket/test/server-slice"
)

var Once sync.Once

func BenchmarkClient(b *testing.B) {
	Once.Do(func() {
		go serverslice.RunServer()
		time.Sleep(1 * time.Second)
	})

	for i := 0; i < b.N; i++ {
		connect := Connection.New("localhost", "8585")
		conn, err := connect.ConnectionClient()

		if err != nil {
			panic(err)
		}

		defer conn.Close()

		people := []global.Custom{
			{Name: "Paulo", Age: 17},
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
}
