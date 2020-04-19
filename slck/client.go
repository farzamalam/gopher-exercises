package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"strconv"
)

// DELIMITER defines the boundary of the message
var (
	DELIMITER = []byte(`\r\n`)
)

// a client is a wrapper around TCP connection. It encaptulates all the functionality around accepting
// the message from TCP connection, parsing the messages, validating their structures and contents
// and sending them to hub for further processing.
// It contains five attributes.
// conn : it is the TCP connection itself.
// outbound : it is send only channel of type command, it is used to send the command to hub for further processing.
// registrer : it is send only channel of type client, it is used to register the client with the hub.
// deregister : it is send only channel of type client, it is used to un-register the client with the hub.
type client struct {
	conn       net.Conn
	outbound   chan<- command
	register   chan<- *client
	deregister chan<- *client
	username   string
}

// newClient returns a new client instance.
func newClient(conn net.Conn, o chan<- command, r chan<- *client, d chan<- *client) *client {
	return &client{
		conn:       conn,
		outbound:   o,
		register:   r,
		deregister: d,
	}
}

// read is used to listen to incoming TCP messages, if recieves the io.EOF then it sends deregister to
// hub to delete the account from all users and channels, or it sends the msg to handle.
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
}

// handle gets the raw message from the socket and parses the bytes to make meaning out of them.
// it splits the raw message into two parts cmd and args, and matches the command with the protocol
// commands and if matches with any then it sends the client and args to that method, otherwise calls err method.
func (c *client) handle(message []byte) {
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

// reg is used to register a new client.
// it validates to see if it has right length and @ in the begining of the username.
// it sends the client itself through the register channel, this channel is read by the hub.
func (c *client) reg(args []byte) error {
	u := bytes.TrimSpace(args)
	if u[0] != '@' {
		return fmt.Errorf("Username must begin with @")
	}
	if len(u) == 0 {
		return fmt.Errorf("Username cannot be blank")
	}
	c.username = string(u)
	c.register <- c
	return nil
}

// join validates the channel name and sends outbound command to hub using chan.
func (c *client) join(args []byte) error {
	channelID := bytes.TrimSpace(args)
	if channelID[0] != '#' {
		return fmt.Errorf("Channel ID must begin with #")
	}
	c.outbound <- command{
		recipient: string(channelID),
		sender:    c.username,
		id:        JOIN,
	}
	return nil
}

// leave validates the channel name and sends the outbound command to hub using chan.
func (c *client) leave(args []byte) error {
	channelID := bytes.TrimSpace(args)
	if channelID[0] != '#' {
		return fmt.Errorf("Channel ID must begin with #")
	}
	c.outbound <- command{
		recipient: string(channelID),
		sender:    c.username,
		id:        LEAVE,
	}
	return nil
}

// msg takes in the args and validates it and sends command via outbound for further processing.
// it takes body, username and recipient along with MSG ID and initializes a command and sends it to hub
// using c.outbound channel.
func (c *client) msg(args []byte) error {
	args = bytes.TrimSpace(args)
	if args[0] != '#' && args[0] != '@' {
		return fmt.Errorf("Recipient must be a channel(#) or user(@)")
	}
	recipient := bytes.Split(args, []byte(" "))[0]
	if len(recipient) == 0 {
		return fmt.Errorf("Recipient must have a name")
	}
	args = bytes.TrimSpace(bytes.TrimPrefix(args, recipient))
	l := bytes.Split(args, DELIMITER)[0]
	length, err := strconv.Atoi(string(l))
	if err != nil {
		return fmt.Errorf("Body length must be present")
	}
	if length == 0 {
		return fmt.Errorf("Body lenght must be atleast 1")
	}
	padding := len(l) + len(DELIMITER)
	body := args[padding : padding+length]
	c.outbound <- command{
		recipient: string(recipient),
		sender:    c.username,
		body:      body,
		id:        MSG,
	}
	return nil
}

// chns sends the CHNS command along with username to hub via outbound channel.
func (c *client) chns() {
	c.outbound <- command{
		sender: c.username,
		id:     CHNS,
	}
}

// usrs sends the USRS command id along with username of sender to hub via outbound channel
func (c *client) usrs() {
	c.outbound <- command{
		sender: c.username,
		id:     USRS,
	}
}

// err is called to print the ERR on the conn, it will not call the hub. Mostly it is used to check the message
// integrity.
func (c *client) err(e error) {
	c.conn.Write([]byte("ERR : " + e.Error() + "\n"))
}
