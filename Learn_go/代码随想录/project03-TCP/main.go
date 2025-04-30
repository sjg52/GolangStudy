package main

import (
	"fmt"
	"net"
)

// main starts a TCP echo server that listens on localhost:8080 and handles incoming client connections.
func main() {
	// Listen for incoming connections
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Ensure we teardown the server when the program exits
	defer listener.Close()

	fmt.Println("Server is listening on port 8080")

	for {
		// Block until we receive an incoming connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		// Handle client connection
		handleClient(conn)
	}
}

// handleClient reads data from a TCP client connection and echoes it back to the client.
func handleClient(conn net.Conn) {
	// Ensure we close the connection after we're done
	defer conn.Close()

	// Read data
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		return
	}

	fmt.Println("Received data", buf[:n])

	// Write the same data back
	conn.Write(buf[:n])
}
