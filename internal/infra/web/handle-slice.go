package web

import (
	"encoding/gob"
	"fmt"
	"net"

	"github.com/rafaelsouzaribeiro/go-socket/pkg/global"
)

func (h *Iconnection) HandleSliceConnection(conn *net.TCPConn, channel chan global.Custom) {
	defer conn.Close()

	var person []global.Custom
	decoder := gob.NewDecoder(conn)
	err := decoder.Decode(&person)
	if err != nil {
		fmt.Println("Error decoding message:", err)
		return
	}

	for _, p := range person {
		channel <- p

	}

}
