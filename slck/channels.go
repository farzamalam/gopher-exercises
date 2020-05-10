package main

// channel type is used to define a chat room.
// it has two attributes,
// name : it defines the unique name of channel.
// clients : it is map of all the clients that present at current inside the channel.
type channel struct {
	name    string
	clients map[*client]bool
}

// newChannel is used to return a new channel. It is used to create a new channel.
func newChannel(name string) *channel {
	return &channel{
		name:    name,
		clients: make(map[*client]bool),
	}
}

// broadcast is used to used messages to all the clients that are present in the channel.
// it takes the username of the sender and the message to send to the channel.
// the output is : `@a:"Hello, How are you?"` to all the member of channel
// the message that is sent to all the clients are in bytes.
func (c *channel) broadcast(s string, m []byte) {
	msg := append([]byte(s), ": "...)
	msg = append(msg, m...)
	msg = append(msg, '\n')

	for cl := range c.clients {
		cl.conn.Write(msg)
	}
}
