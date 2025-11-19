package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", "localhost:4666")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Println("Running on port 4666\nwaiting for client....")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
}
