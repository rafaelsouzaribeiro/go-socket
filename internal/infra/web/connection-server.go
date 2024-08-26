package web

import (
	"fmt"
	"net"
)

func (h *Iconnection) ConnectionServer() (*net.TCPListener, error) {

	connStr := fmt.Sprintf("%s:%s", h.Host, h.Port)
	address, _ := net.ResolveTCPAddr("tcp", connStr)

	listener, err := net.ListenTCP("tcp", address)

	if err != nil {
		return nil, err
	}

	return listener, nil
}
