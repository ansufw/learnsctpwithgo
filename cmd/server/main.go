package main

import (
	"fmt"
	"net"

	"github.com/ishidawataru/sctp"
)

func main() {
	// Define multiple IP addresses for the SCTP server (multihomed)
	addr := &sctp.SCTPAddr{
		IPAddrs: []net.IPAddr{
			{IP: net.ParseIP("192.168.1.100")},
			{IP: net.ParseIP("10.0.0.1")},
		},
		Port: 8080,
	}

	// Listen on the multihomed addresses
	listener, err := sctp.ListenSCTP("sctp", addr)
	if err != nil {
		fmt.Println("Error starting SCTP server:", err)
		return
	}
	defer listener.Close()
	fmt.Println("SCTP server started on multiple addresses")

	for {
		conn, err := listener.AcceptSCTP()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn *sctp.SCTPConn) {
	defer conn.Close()
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			return
		}
		fmt.Printf("Received: %s\n", string(buffer[:n]))
		conn.Write([]byte("Message received via SCTP\n"))
	}
}
