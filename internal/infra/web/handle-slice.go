package web

import (
	"encoding/gob"
	"fmt"
	"net"
)

func (h *Iconnection) HandleSliceConnection(conn *net.TCPConn) {
	defer conn.Close()

	var person []Person
	decoder := gob.NewDecoder(conn)
	err := decoder.Decode(&person)
	if err != nil {
		fmt.Println("Error decoding message:", err)
		return
	}

	for _, p := range person {
		fmt.Printf("Message received. Name: %s, Age: %d\n", p.Name, p.Age)
	}

}
