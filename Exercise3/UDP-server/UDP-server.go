package main

import (
	"fmt"
	"math/rand"
	"net"
	"strings"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	s, err := net.ResolveUDPAddr("udp", ":20008")
	if err != nil {
		fmt.Println(err)
		return
	}

	connection, err := net.ListenUDP("udp", s)
	if err != nil {
		fmt.Println(err)
		returncd
	}

	//defer connection.Close()
	buffer := make([]byte, 1024)

	rand.Seed(time.Now().Unix())

	for {
		n, addr, err := connection.ReadFromUDP(buffer)
		fmt.Print("-> ", string(buffer[0:n-1]))

		if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
			fmt.Println("Exiting UDP server!")
			return
		}

		data := []byte("this is a message from space")
		fmt.Printf("data: %s\n", string(data))
		//ser, err := net.ResolveUDPAddr("udp", "10.100.23.187:20008")
		_, err = connection.WriteToUDP(data, addr)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
