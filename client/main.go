package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const (
	serverHost = "localhost"
	serverPort = "5000"
)

func main() {
	fmt.Println("Connecting to server ", serverHost + ":" + serverPort)

	conn, err := net.Dial("tcp", serverHost + ":" + serverPort)

	if err != nil {
		fmt.Println("Error connecting: ", err.Error())
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Text to send: ")

		input, _ := reader.ReadString('\n')

		if _, err := conn.Write([]byte(input)); err != nil {
			fmt.Println("Error while writing to the socket: ", err.Error())
			continue;
		}

		message, _ := bufio.NewReader(conn).ReadString('\n')

		log.Print("[" + serverHost + ":" + serverPort + "] : " + message)
	}
}
