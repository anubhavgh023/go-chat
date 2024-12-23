package main

import (
	"fmt"
	"log"
	"net"
)

const PORT = "8080"

func handleConnection(conn net.Conn) {
	defer conn.Close()
	msg := fmt.Sprintf("Welcome Soldier !\n")
	n, err := conn.Write([]byte(msg))
	if err != nil {
		log.Printf("ERROR: could not write to %s: %s\n", conn.RemoteAddr().String(), err)
		return
	}

	if n < len(msg) {
		log.Printf("The message was not fully written %d/%d\n", n, len(msg))
	}
}

func main() {
	ln, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("ERROR: could not listen to port %s: %s\n", PORT, err)
	}
	fmt.Println("Server is running on PORT:", 8080)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("ERROR: cound not accept a connection: %s\n", err)
		}
		log.Printf("Accepted connection from %s:", conn.RemoteAddr().String())
		go handleConnection(conn)
	}
}
