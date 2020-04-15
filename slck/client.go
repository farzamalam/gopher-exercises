package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
)

var (
	DELIMITER = []byte(`\r\n`)
)

type client struct {
	conn       net.Conn
	outbound   chan<- command
	register   chan<- *client
	deregister chan<- *client
	username   string
}

func newClient(conn net.Conn, o chan<- command, r chan<- *client, d chan<- *client) *client {
	return &client{
		conn:       conn,
		outbound:   o,
		register:   r,
		deregister: d,
	}
}

func (c *client) read() error {
	for {
		msg, err := bufio.NewReader(c.conn).ReadBytes('\n')
		if err == io.EOF {
			c.deregister <- c
			return nil
		}
		if err != nil {
			return err
		}
		c.handle(msg)
	}
	return nil
}

func handle(message []byte) {
	cmd := bytes.ToUpper(bytes.TrimSpace(bytes.Split(message, []byte(" "))[0]))
	args := bytes.TrimSpace(bytes.TrimPrefix(message, cmd))
	switch string(cmd) {
	case "REG":
		if err := c.reg(args); err != nil {
			c.err(err)
		}
	case "JOIN":
		if err := c.join(args); err != nil {
			c.err(err)
		}
	case "LEAVE":
		if err := c.leave(args); err != nil {
			c.err(err)
		}
	case "MSG":
		if err := c.msg(args); err != nil {
			c.err(err)
		}
	case "CHNS":
		c.chns()
	case "USRS":
		c.usrs()
	default:
		c.err(fmt.Errorf("Unknown command %s", cmd))
	}
}
