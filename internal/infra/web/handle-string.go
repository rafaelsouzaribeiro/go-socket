package web

import (
	"fmt"
	"net"
)

func (h *Iconnection) HandleStringConnection(conn *net.TCPConn, channel chan<- string) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading from connection:", err)
		return
	}

	channel <- string(buffer[:n])

}
