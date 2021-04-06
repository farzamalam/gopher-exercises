package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"runtime"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/sync/semaphore"
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

	sem := semaphore.NewWeighted(int64(numWorkers))
	ctx := context.TODO()
	var openPorts []int

	for _, port := range portsToScan {
		if err := sem.Acquire(ctx, 1); err != nil {
			log.Printf("Failed to acquire semaphore: %v\n", err)
			break
		}
		go func(port int) {
			defer sem.Release(1)
			p := scan(host, port)
			if p > 0 {
				openPorts = append(openPorts, p)
			}
		}(port)
	}

	err = sem.Acquire(ctx, int64(numWorkers))
	if err != nil {
		fmt.Printf("Failed to acquire semaphore: %v\n", err)
	}

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

func scan(host string, port int) int {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Printf("%d CLOSED\n", port)
		return -1
	}
	conn.Close()
	return port
}
