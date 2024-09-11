package main

import (
	"fmt"
	"net"

	"github.com/ishidawataru/sctp"
)

func main() {
	addr := &sctp.SCTPAddr{
		IPAddrs: []net.IPAddr{
			{IP: net.ParseIP("192.168.1.100")},
			{IP: net.ParseIP("10.0.0.1")},
		},
		Port: 8080,
	}

	// Create an SCTP client with multiple IP addresses (multihomed)
	conn, err := sctp.DialSCTP("sctp", nil, addr)
	if err != nil {
		fmt.Println("Error connecting to SCTP server:", err)
		return
	}
	defer conn.Close()

	message := "Hello, SCTP Multihoming Server!"
	conn.Write([]byte(message))
	fmt.Println("Sent message to server:", message)

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading from server:", err)
		return
	}
	fmt.Printf("Server reply: %s\n", string(buffer[:n]))
}
