package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"runtime"
	"sort"
	"strconv"
	"strings"
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
	portsToScan, err := parsePortsToScan(ports)
	if err != nil {
		log.Fatalf("Unable to parse ports flag : %v", err)
	}

	portsChan := make(chan int, numWorkers)
	resultsChan := make(chan int)

	for i := 0; i < cap(portsChan); i++ {
		go worker(host, portsChan, resultsChan)
	}

	go func() {
		for _, p := range portsToScan {
			portsChan <- p
		}
	}()

	var openPorts []int
	for i := 0; i < len(portsToScan); i++ {
		if p := <-resultsChan; p != -1 {
			openPorts = append(openPorts, p)
		}
	}
	close(portsChan)
	close(resultsChan)

	sort.Ints(openPorts)
	for _, p := range openPorts {
		log.Printf("%d OPEN \n", p)
	}
}

func parsePortsToScan(portsFlag string) ([]int, error) {
	p, err := strconv.Atoi(portsFlag)
	if err == nil {
		return []int{p}, nil
	}
	ports := strings.Split(portsFlag, "-")
	if len(ports) != 2 {
		return nil, errors.New("unable to determine port(s) to scan")
	}

	fp, err := strconv.Atoi(ports[0])
	if err != nil {
		return nil, fmt.Errorf("falied to convert %s to a valid port number", ports[0])
	}
	tp, err := strconv.Atoi(ports[1])
	if err != nil {
		return nil, fmt.Errorf("failed to convert %s to a valid port number", ports[1])
	}
	if tp < 0 || fp < 0 {
		return nil, fmt.Errorf("port number must be greater than 0")
	}
	var res []int
	for p := fp; p <= tp; p++ {
		res = append(res, p)
	}
	return res, nil
}

func worker(host string, portsChan <-chan int, resultsChan chan<- int) {
	for p := range portsChan {
		address := fmt.Sprintf("%s:%d", host, p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			fmt.Printf("%d CLOSED %s\n", p, err)
			resultsChan <- -1
			continue
		}
		conn.Close()
		resultsChan <- p
	}
}
