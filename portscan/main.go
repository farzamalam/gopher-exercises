package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

var host string
var ports string
var numWorkers int

func init() {
	flag.StringVar(&host, "host", "localhost", "host to scan.")
	flag.StringVar(&ports, "ports", "80", "Port(s) ex 80, 443, 8080-9080.")
	flag.IntVar(&numWorkers, "workers", runtime.NumCPU(), "Number of workers. Defaults to 8.")
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

func portsToScan(portsFlag string) ([]int, error) {
	p, err := strconv.Atoi(portsFlag)
	if err == nil {
		return []int{p}, nil
	}
	ports := strings.Split(portsFlag, "-")
	if len(ports) != 2 {
		return nil, errors.New("unable to determine port(s) to scan.")
	}

	fp, err := strconv.Atoi(ports[0])
	if err != nil {
		return nil, fmt.Errorf("Falied to convert %s to a valid port number.", ports[0])
	}
	tp, err := strconv.Atoi(ports[1])
	if err != nil {
		return nil, fmt.Errorf("Failed to convert %s to a valid port number.", ports[1])
	}
	if tp < 0 || fp < 0 {
		return nil, fmt.Errorf("Port number must be greater than 0.")
	}
	var res []int
	for p := fp; p <= tp; p++ {
		res = append(res, p)
	}
	return res, nil
}
