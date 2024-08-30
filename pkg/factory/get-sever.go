package factory

import (
	"net"

	Connection "github.com/rafaelsouzaribeiro/go-socket/internal/infra/web"
)

func (f *FactoryServer) GetServer(connect *Connection.Iconnection, conn *net.TCPConn) {
	if f.types == Slice {
		connect.HandleStructConnection(conn, f.Channel)
	}

	if f.types == Struct {
		connect.HandleSliceConnection(conn, f.Channel)
	}
}
