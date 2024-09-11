package web

import (
	"encoding/binary"
	"net"
)

func (h *Iconnection) HandleIntConnection(conn *net.TCPConn, channel chan<- int32) {
	defer conn.Close()

	var num int32
	err := binary.Read(conn, binary.LittleEndian, &num)
	if err != nil {
		panic(err)
	}

	channel <- num

}
