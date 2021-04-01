package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	h := "localhost"
	p := 8081
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", h, p))
	if err != nil {
		log.Fatalf("PORT %d is CLOSED : %v\n", p, err)
	}
	conn.Close()
	log.Printf("PORT %d is OPEN\n", p)
}
