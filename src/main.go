package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
)

func main() {
	host := flag.String("ip", "google.com", "Target IP address or hostname")
	port := flag.Int("p", 80, "TCP port to connect to (1-65535)")
	count := flag.Int("c", 3, "Number of connection attempts")
	verbose := flag.Bool("vb", false, "Enable detailed connection logging")
	flag.Parse()
	if *port < 1 || *port > 65535 {
		fmt.Println("error: port must be between 1 and 65535")
		return
	}
	if *count <= 0 {
		fmt.Println("error: count must be greater than 0")
		return
	}
	address := net.JoinHostPort(*host, strconv.Itoa(*port))
	startping(address, *count, *verbose)
}
