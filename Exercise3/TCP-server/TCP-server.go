package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", "10.100.23.147:33546")
	if err != nil {
		log.Fatal("tcp server listener error:", err)
	}

	for {
		conn, _ := ln.Accept()
		// message, _ := bufio.NewReader(conn).ReadString('\n')
		// fmt.Print("Message Received: " + string(message))
		// conn.Write([]byte(strings.ToUpper(message) + "\n"))
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	bufferBytes, err := bufio.NewReader(conn).ReadBytes('\n')

	if err != nil {
		log.Println("Client left... ")
		conn.Close()
		return
	}

	message := string(bufferBytes)
	response := fmt.Sprintf(message + "\n")

	log.Println(response)
	conn.Write([]byte("you sent: " + response))

	handleConnection(conn)

}
