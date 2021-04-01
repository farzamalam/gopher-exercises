package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	h := "localhost"
	fp := 8075
	tp := 8085
	for p := fp; p <= tp; p++ {
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", h, p))
		if err != nil {
			log.Printf("%d - CLOSED : %v\n", p, err)
			continue
		}
		conn.Close()
		log.Printf("%d - OPEN\n", p)
	}

}
