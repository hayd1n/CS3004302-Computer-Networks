package iperfer

import (
	"errors"
	"fmt"
	"io"
	"net"
	"time"
)

type Server struct {
	listener net.Listener
}

func NewServer(port uint16) (*Server, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		return nil, errors.New("initialize new server failed")
	}

	return &Server{listener: listener}, nil
}

func (c *Server) Receive() (uint64, float32, error) {
	conn, err := c.listener.Accept()
	if err != nil {
		return 0, 0.0, errors.New("connection failed")
	}
	defer conn.Close()

	start_time := time.Now()
	n, err := handle(conn)
	if err != nil {
		return n, 0.0, errors.New("handling connection failed")
	}
	bandwidth := calcBandwidth(n, int(time.Since(start_time).Seconds()))

	return n, bandwidth, nil
}

func handle(conn net.Conn) (uint64, error) {
	var totalReceived uint64
	buffer := make([]byte, 65535)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			// handle the error, check if it's a normal connection close
			if err == io.EOF {
				return totalReceived, nil
			}

			return totalReceived, errors.New("reading data failed")
		}

		totalReceived += uint64(n)
	}
}
