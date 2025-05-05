package main

import (
	"fmt"
	"log"
	"net"

	"github.com/codecrafters-io/dns-server-starter-go/internal/server"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:2053")
	if err != nil {
		fmt.Println("Failed to resolve UDP address:", err)
		return err
	}

	return server.NewDnsServer(udpAddr).Serve()
}
