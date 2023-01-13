package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:9999")
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := listener.Close(); err != nil {
			log.Println("failed to close TCP listener:", err.Error())
		}
	}()
	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Println("failed to accept connection:", err)
		}
		go handleConnection(&connection)
	}
}

func handleConnection(connection *net.Conn) {
	log.Println("incoming connection...")
	for {
		data := make([]byte, 4096)
		_, err := (*connection).Read(data)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println("failed to read incoming data:", err)
		}
		fmt.Println(string(data))
	}
	log.Println("connection closing...")
	err := (*connection).Close()
	if err != nil {
		log.Fatalln("failed to read incoming data:", err)
	}
}
