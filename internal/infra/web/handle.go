package web

import (
	"encoding/gob"
	"fmt"
	"net"
)

func (h *Iconnection) HandleConnection(conn *net.TCPConn) {
	defer conn.Close()

	var send Send
	decoder := gob.NewDecoder(conn)
	err := decoder.Decode(&send)
	if err != nil {
		fmt.Println("Error decoding message:", err)
		return
	}

	fmt.Println("Message received:", send.Message)

}
