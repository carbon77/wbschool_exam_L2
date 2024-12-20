package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// TCP сервер для тестирования

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		inputMessage, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Failed to read message: %v\n", err)
		}
		fmt.Printf("Received message: %s\n", inputMessage)

		outputMessage := fmt.Sprintf("localhost:8080> %s", inputMessage)
		if _, err := conn.Write([]byte(outputMessage)); err != nil {
			log.Fatalf("Failed to write message: %v\n", err)
		}
	}
}

func main() {

	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v\n", err)
	}
	defer listener.Close()

	fmt.Println("Listening on localhost:8080")

	for {
		conn, err := listener.Accept()
		fmt.Printf("Accepted connection from %s\n", conn.RemoteAddr())
		if err != nil {
			log.Fatalf("Failed to accept connection: %v\n", err)
		}

		go handleConnection(conn)
	}
}
