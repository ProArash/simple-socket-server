package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Printf("Client connected: %s\n", conn.RemoteAddr().String())
	welcomeMsg := fmt.Sprintf("Welcome %s\n", conn.RemoteAddr().String())
	conn.Write([]byte(welcomeMsg))
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Client %s disconnected: %s\n", conn.RemoteAddr().String(), err.Error())
			break
		}
		trimmedMessage := message[:len(message)-1]
		fmt.Printf("%s From %s\n", trimmedMessage, conn.RemoteAddr().String())
		msg := fmt.Sprintf("[%s]: %s", conn.RemoteAddr().String(), message)
		conn.Write([]byte(msg))
		if trimmedMessage == "STOP" {
			fmt.Printf("Closed connection by: %s\n", conn.RemoteAddr().String())
			break
		}
	}
}
