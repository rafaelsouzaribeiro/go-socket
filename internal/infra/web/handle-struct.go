package web

import (
	"encoding/gob"
	"fmt"
	"net"
)

func (h *Iconnection) HandleStructConnection(conn *net.TCPConn, channel chan Person) {
	defer conn.Close()

	var person Person
	decoder := gob.NewDecoder(conn)
	err := decoder.Decode(&person)
	if err != nil {
		fmt.Println("Error decoding message:", err)
		return
	}

	channel <- person

}
