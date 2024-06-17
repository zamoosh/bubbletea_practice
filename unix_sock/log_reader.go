package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	const socketPath = "/W/code_and_project/code_project/go/bubbletea_practice/unix_sock/socket.sock"

	if err := os.RemoveAll(socketPath); err != nil {
		log.Fatalf("Failed to remove existing socket: %v", err)
	}

	addr, err := net.ResolveUnixAddr("unix", socketPath)
	if err != nil {
		log.Fatalf("Failed to resolve Unix address: %v", err)
	}

	listener, err := net.ListenUnix("unix", addr)
	if err != nil {
		log.Fatalf("Failed to listen on Unix socket: %v", err)
	}
	defer listener.Close()

	fmt.Println("Log reader started, waiting for logs...")

	for {
		conn, err := listener.AcceptUnix()
		if err != nil {
			log.Fatalf("Failed to accept connection: %v", err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn *net.UnixConn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Printf("Failed to read from connection: %v", err)
			return
		}

		fmt.Printf("Received log message: %s\n", string(buffer[:n]))
	}
}
