package web

import (
	"fmt"
	"net"
	"time"
)

func (h *Iconnection) HandleConnection(conn *net.TCPConn) {
	defer conn.Close()

	conn.SetDeadline(time.Now().Add(10 * time.Second))
	item := make([]byte, 1024)
	size, err := conn.Read(item)
	if err != nil {
		fmt.Println("Error reading data:", err)
		return
	}

	message := string(item[:size])
	fmt.Println("Received:", message)

}
