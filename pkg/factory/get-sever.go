package factory

import (
	"net"

	Connection "github.com/rafaelsouzaribeiro/go-socket/internal/infra/web"
)

func (f *FactoryServer) GetServer(connect *Connection.Iconnection, conn *net.TCPConn) {
	if f.types == Struct {
		connect.HandleStructConnection(conn, f.ChannelPerson)
	}

	if f.types == Slice {
		connect.HandleSliceConnection(conn, f.ChannelPerson)
	}

	if f.types == String {
		connect.HandleStringConnection(conn, f.ChannelString)
	}
}
