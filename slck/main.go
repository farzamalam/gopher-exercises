package main

import (
	"log"
	"net"
)

<<<<<<< HEAD
=======
// main is used to initalize a listen to tcp connection and starts hub and read.
>>>>>>> a7c690b89df0c0f865afc74ab64c6c57aa3897b0
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
<<<<<<< HEAD
		c := newClient(
			conn,
			hub.commands,
			hub.registration,
			hub.deregistration,
		)
=======
		c := newClient(conn, hub.commands, hub.registration, hub.deregistration)
>>>>>>> a7c690b89df0c0f865afc74ab64c6c57aa3897b0
		go c.read()
	}
}
