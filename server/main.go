package main

import (
	"fmt"
	"log"
	"net"
)

func process(conn net.Conn) {
	// read the message form the client
	defer conn.Close()

	for {
		buffer := make([]byte, 8096)
		n, err := conn.Read(buffer[:4])
		if n != 4 || err != nil {
			log.Println("conn.Read message error", err)
			return
		}
		log.Println("Reading buffer=", buffer[:4])
	}
}

func main() {

	fmt.Println("Server is listening port:8889....")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		log.Println("net.Listen err=", err)
		return
	}
	defer listen.Close()

	// Once listening success, waiting for the client to connect server

	for {
		fmt.Println("Waiting clients to connect the server.....")
		conn, err := listen.Accept()
		if err != nil {
			log.Println("listen Accept err=", err)
		}

		// Once connection success, open a goroutine and keep communicating with the client

		go process(conn)
	}
}
