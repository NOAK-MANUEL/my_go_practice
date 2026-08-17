package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/spf13/pflag"
)

func checkTCP(port string, host string) Result {
	result := Result{
		port: port,
	}
	conn, err := net.DialTimeout("tcp", host+":"+port, 1000*time.Millisecond)

	if err != nil {

		return result
	}
	defer conn.Close()
	result.open = true
	return result
}

type Result struct {
	port string
	open bool
}

func main() {
	var ports string
	pflag.StringVarP(&ports, "ports", "p", "3000", "Port range e.g 10-3000 or 3000 or 8080")
	pflag.Parse()
	args := pflag.Args()
	if len(args) < 1 {
		log.Fatal("Incomplete argument <host> <port>")
	}
	results := make(chan Result)
	var wg sync.WaitGroup
	host := args[0]
	var total = 0
	if strings.Contains(ports, "-") {
		input := strings.Split(ports, "-")
		start, err := strconv.Atoi(input[0])
		if err != nil {
			log.Fatal("Not a valid port", err)
		}
		end, err := strconv.Atoi(input[1])
		if err != nil {
			log.Fatal("Not a valid port", err)
		}
		total = end - start + 1
		for port := start; port <= end; port++ {
			wg.Add(1)
			go func(port int) {
				defer wg.Done()
				results <- checkTCP(strconv.Itoa(port), host)

			}(port)
		}

	} else {
		portSlice := strings.Split(ports, ",")
		total = len(portSlice)
		for _, port := range portSlice {
			wg.Add(1)
			go func(port string) {
				defer wg.Done()
				results <- checkTCP(port, host)
			}(port)
		}
	}

	for i := 0; i < total; i++ {
		result := <-results
		if result.open {
			fmt.Printf("Port %s connected successfully\n", result.port)
		} else {
			fmt.Printf("Port %s unsuccessfully connected\n", result.port)
		}
	}

	wg.Wait()
}
