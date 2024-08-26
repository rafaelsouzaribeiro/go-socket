package web

import (
	"fmt"
	"net"
)

func (h *Iconnection) ConnectionClient() (net.Conn, error) {

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", h.Host, h.Port))
	if err != nil {
		return nil, err
	}

	return conn, nil
}
