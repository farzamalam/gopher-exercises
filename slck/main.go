package main

import (
	"log"
	"net"
)

// main is used to initalize a listen to tcp connection and starts hub and read.
func main() {
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Printf("%v", err)
	}
	hub := newHub()
	go hub.run()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("%v", err)
		}
		c := newClient(conn, hub.commands, hub.registration, hub.deregistration)
		go c.read()
	}
}
