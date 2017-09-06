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
	conn, err := net.Dial(TYPE, ADDRESS)

	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}

	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, "%s\n", text)

		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("Message from server:", message)
	}
}
