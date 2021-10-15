package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

const (
	host = "localhost"
	port         = 5000
)

func main() {
	startMessage := fmt.Sprintf("Listening on %s:%d", host, port)
	fmt.Println(startMessage)

	listener, err := net.Listen("tcp", host + ":" + strconv.Itoa(port))
	if err != nil {
		fmt.Println("Error while listening: ", err.Error())
		os.Exit(1)
	}

	defer func() {
		if err := listener.Close(); err != nil {
			fmt.Println("Error while closing listener: ", err.Error())
			os.Exit(1)
		}
	}()

	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println("Error while connecting: ", err.Error())
			return
		}

		connectionMessage := fmt.Sprintf("Client from %s connected.", connection.RemoteAddr().String())
		fmt.Println(connectionMessage)

		go handleClientConnection(connection)
	}
}

func handleClientConnection(connection net.Conn) {
	buffer, err := bufio.NewReader(connection).ReadBytes('\n')
	if err != nil {
		connectionClosedMessage := fmt.Sprintf("Client from %s left.", connection.RemoteAddr().String())
		fmt.Println(connectionClosedMessage);

		if err := connection.Close(); err != nil {
			fmt.Println("Error while closing connection: ", err.Error())
		}

		return
	}

	connectionMessage := fmt.Sprintf("[%s] : %s", connection.RemoteAddr().String(), string(buffer[:len(buffer) - 1]))
	log.Println(connectionMessage)

	if _, err:= connection.Write(buffer); err != nil {
		fmt.Println("Error while writing to socket with " + connection.RemoteAddr().String() + " : ", err.Error())
	}

	handleClientConnection(connection)
}
