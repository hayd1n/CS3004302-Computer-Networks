package iperfer

import (
	"errors"
	"fmt"
	"net"
	"time"
)

type Client struct {
	conn net.Conn
}

func NewClient(host string, port uint16) (*Client, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return nil, errors.New("initialize new client failed")
	}

	return &Client{conn: conn}, nil
}

func (c *Client) Send(seconds int) (uint64, float32, error) {
	// generate data chunk
	var sendData []byte
	for i := 0; i < 1000; i++ {
		sendData = append(sendData, 0)
	}

	start_time := time.Now()
	n := uint64(0)
	duration := time.Second * time.Duration(seconds)
	for time.Since(start_time) < duration {
		wn, err := c.conn.Write(sendData)
		if err != nil {
			fmt.Println(err)
			return n, 0.0, errors.New("sending data failed")
		}

		n += uint64(wn)
	}

	bandwidth := calcBandwidth(n, seconds)

	return n, bandwidth, nil
}
