package factory

import (
	"net"

	Connection "github.com/rafaelsouzaribeiro/go-socket/internal/infra/web"
)

func (f *FactoryServer) GetServer(connect *Connection.Iconnection, conn *net.TCPConn) {
	if f.typeStruct {
		connect.HandleStructConnection(conn, f.Channel)
	}

	if f.typeSlice {
		connect.HandleSliceConnection(conn, f.Channel)
	}
}
