package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1)
	}
}

func main() {
	// if len(os.Args) != 3 {
	// 	fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
	// 	os.Exit(1)
	// }
	// service := os.Args[1]
	service := "10.100.23.147:20008"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)
	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)
	lisAddr, err := net.ResolveUDPAddr("udp4", ":20008")
	checkError(err)
	lisconn, err := net.ListenUDP("udp", lisAddr)
	checkError(err)

	var buf [512]byte
	go func() {
		for {
			_, err = conn.Write([]byte("suh dude \n"))
			time.Sleep(5 * time.Second)
		}
	}()
	go func() {
		for {
			n, err := lisconn.Read(buf[0:])
			checkError(err)
			fmt.Println(string(buf[0:n]))
		}
	}()
	for {

	}
}
