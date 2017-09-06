package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	ADDRESS = "localhost:6789"
	TYPE    = "tcp"
)

func main() {
	net, err := net.Listen(TYPE, ADDRESS)

	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	fmt.Println("Listening on:", ADDRESS)

	for {
		conn, err := net.Accept()

		if err != nil {
			fmt.Println("Error accepting:", err.Error())
			os.Exit(1)
		}

		go handleRequest(conn)
	}

	defer net.Close()
}

func handleRequest(conn net.Conn) {
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			fmt.Println("Error reading:", err.Error())
			os.Exit(1)
		}

		fmt.Print("Message received: ", message)
		conn.Write([]byte(message))
	}

	defer conn.Close()
}
