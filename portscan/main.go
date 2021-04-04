package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
)

var host string
var fromPort string
var toPort string

func init() {
	flag.StringVar(&host, "host", "localhost", "host to scan.")
	flag.StringVar(&fromPort, "from", "8080", "Port to start scanning from.")
	flag.StringVar(&toPort, "to", "8090", "Port to end scanning to.")
}

func main() {
	flag.Parse()
	fp, err := strconv.Atoi(fromPort)
	if err != nil {
		log.Fatalf("Error while parsing 'from' port: %v", err)
	}
	tp, err := strconv.Atoi(toPort)
	if err != nil {
		log.Fatalf("Error while parsing 'to' port : %v", err)
	}

	if fp > tp {
		log.Fatalf("Invalid values of 'from' and 'to' port.")
	}

	var wg sync.WaitGroup

	wg.Add(tp - fp + 1)
	for p := fp; p <= tp; p++ {
		go func(p int) {
			defer wg.Done()
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, p))
			if err != nil {
				log.Printf("%d Closed %s\n", p, err)
				return
			}
			conn.Close()
			log.Printf("%d Open\n", p)
		}(p)
	}
	log.Println("Waiting")
	wg.Wait()
	log.Println("Done.")

}
