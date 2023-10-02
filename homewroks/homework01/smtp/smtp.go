package smtp

import (
	"errors"
	"fmt"
	"net"
)

var (
	SendTCPError = errors.New("failed to sendTCP to the server")
)

type Smtp struct {
	conn net.Conn
	Host string
}

func NewSmtp() *Smtp {
	return &Smtp{}
}

func (c *Smtp) Connect() error {
	conn, err := net.Dial("tcp", c.Host)
	if err != nil {
		return errors.New("failed to connect to the server")
	}

	c.conn = conn
	return nil
}

func (c *Smtp) Send(from string, to string, subject string, body string) error {
	if _, err := c.sendTCP("HELO", false); err != nil {
		return SendTCPError
	}

	if _, err := c.sendTCP(fmt.Sprintf("MAIL FROM: <%s>", from), false); err != nil {
		return SendTCPError
	}

	if _, err := c.sendTCP(fmt.Sprintf("RCPT TO: <%s>", to), false); err != nil {
		return SendTCPError
	}

	if _, err := c.sendTCP("DATA", false); err != nil {
		return SendTCPError
	}

	if _, err := c.sendTCP(fmt.Sprintf("From: %s", from), false); err != nil {
		return SendTCPError
	}

	if _, err := c.sendTCP(fmt.Sprintf("To: %s", to), false); err != nil {
		return SendTCPError
	}

	if _, err := c.sendTCP(fmt.Sprintf("Subject: %s\r\n", subject), false); err != nil {
		return SendTCPError
	}

	if _, err := c.sendTCP(fmt.Sprintf("%s\r\n.\r\n", body), false); err != nil {
		return SendTCPError
	}

	if _, err := c.sendTCP("QUIT", false); err != nil {
		return SendTCPError
	}

	return nil
}

func (c *Smtp) sendTCP(s string, read_result bool) (string, error) {
	ss := fmt.Sprintf("%s\r\n", s)

	_, err := c.conn.Write([]byte(ss))
	if err != nil {
		return "", err
	}

	if !read_result {
		return "", nil
	}

	var buf [1024]byte
	n, err := c.conn.Read(buf[:])
	if err != nil {
		return "", err
	}

	return string(buf[:n]), nil
}
