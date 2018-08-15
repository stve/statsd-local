package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addr := net.UDPAddr{
		Port: 8125,
		IP:   net.ParseIP("127.0.0.1"),
	}
	conn, err := net.ListenUDP("udp", &addr)
	defer conn.Close()
	if err != nil {
		panic(err)
	}

	fmt.Printf("UDP server running, listening on %d\n", addr.Port)

	buffer := make([]byte, 1024)
	for {
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(buffer[:n]))
	}
}
