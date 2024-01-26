package main

import (
	"log"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", ":1729")
	if err != nil {
		log.Fatal(err)
	}

	for {
		log.Println("waiting for a client to connect")
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("client connected")

		go do(conn)
	}
}

func do(conn net.Conn) {
	// var buf []byte - DOESN'T WORK since it can't write to empty slice
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("input", string(buf))

	_, err = conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello, World!\r\n"))
	if err != nil {
		log.Fatal(err)
	}

	err = conn.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connection closed")
}
