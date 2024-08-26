package web

import (
	"encoding/gob"
	"fmt"
	"net"
)

func (h *Iconnection) HandleConnection(conn *net.TCPConn) {
	defer conn.Close()

	var person Person
	decoder := gob.NewDecoder(conn)
	err := decoder.Decode(&person)
	if err != nil {
		fmt.Println("Error decoding message:", err)
		return
	}

	fmt.Printf("Message received. Name:%s,Age:%d \n", person.Name, person.Age)

}
