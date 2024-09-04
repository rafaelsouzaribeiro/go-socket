package factory

import (
	"net"

	Connection "github.com/rafaelsouzaribeiro/go-socket/internal/infra/web"
)

func (f *FactoryServer) GetServer(connect *Connection.Iconnection, conn *net.TCPConn) {
	if f.types == Struct {
		connect.HandleStructConnection(conn, f.ChannelCustom)
	}

	if f.types == Slice {
		connect.HandleSliceConnection(conn, f.ChannelCustom)
	}

	if f.types == String {
		connect.HandleStringConnection(conn, f.ChannelString)
	}

	if f.types == Int32 {
		connect.HandleIntConnection(conn, f.ChannelInt)
	}

	if f.types == Folder {
		connect.HandleFolderConnection(conn, f.OutputFolderFileName, f.ChannelString)
	}
}
