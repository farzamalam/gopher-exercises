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
		go func(p int) {
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", h, p))
			if err != nil {
				log.Printf("%d - CLOSED : %v\n", p, err)
				return
			}
			conn.Close()
			log.Printf("%d - OPEN\n", p)
		}(p)

	}
	log.Printf("DONE!\n")
}
