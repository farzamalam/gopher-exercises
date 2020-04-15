package main

type ID int

// commands name
const (
	REG ID = iota
	JOIN
	LEAVE
	MSG
	CHNS
	USRS
)

type command struct {
	id        ID
	recipient string
	sendor    string
	body      []byte
}
