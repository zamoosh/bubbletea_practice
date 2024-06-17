package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

var (
	t = time.NewTicker(time.Second * 2)
	c = 10
)

func main() {
	const socketPath = "/W/code_and_project/code_project/go/bubbletea_practice/unix_sock/socket.sock"

	addr, err := net.ResolveUnixAddr("unix", socketPath)
	if err != nil {
		log.Fatalf("Failed to resolve Unix address: %v", err)
	}

	fmt.Println("addr: ", addr)

	conn, err := net.DialUnix("unix", nil, addr)
	if err != nil {
		log.Fatalf("Failed to dial Unix socket: %v", err)
	}
	defer conn.Close()

	var logMessage string
	for c > 0 {
		logMessage = fmt.Sprintf("[%d] %s", c, "This is a log message")
		<-t.C
		if _, err := conn.Write([]byte(logMessage)); err != nil {
			log.Fatalf("Failed to write to Unix socket: %v", err)
		}
		c--
	}
}
