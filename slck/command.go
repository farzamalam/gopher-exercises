package main

// ID type is int, it is used to assign command id values to constants
type ID int

// ID constants, these are used so that we can control the valid command type.
// Using ID we ensure that id(command) will always be a valid command.
const (
	REG ID = iota
	JOIN
	LEAVE
	MSG
	CHNS
	USRS
)

// command(s) are what flows from clients to hub.
// The flow of command will be like, client recieves the the wire-protocol message
// parses it, turns into command and that the client sends to hub.
type command struct {
	id        ID
	recipient string
	sender    string
	body      []byte
}
