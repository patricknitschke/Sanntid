package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	server_addr, err := net.ResolveTCPAddr("tcp", "10.100.23.147:33546")
	conn, err := net.DialTCP("tcp", nil, server_addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = conn.Write([]byte("Connect to: 10.100.23.158:20345" + "\x00"))

	local_addr, err := net.ResolveTCPAddr("tcp", "10.100.23.158:20345")
	conn_ser, err := net.ListenTCP("tcp", local_addr)

	if err != nil {
		fmt.Println("you suck")
		return
	}
	conn_server, err := conn_ser.Accept()
	if err != nil {
		fmt.Println("balls")
		fmt.Println(err)
		return
	}
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n"+"\x00")
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}
}
