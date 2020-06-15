package main

import (
	"log"
	"net"
)

func main() {
	port :="8081"
	fmt.Println("Server started listening on ",port)
	ln, err := net.Listen("tcp", ":"+port)
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
		c := newClient(
			conn,
			hub.commands,
			hub.registration,
			hub.deregistration,
		)
		c = newClient(conn, hub.commands, hub.registration, hub.deregistration)
		go c.read()
	}
}
