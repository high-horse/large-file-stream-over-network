package main

import (
	"fmt"
	"log"
	"net"
)

type FileServer struct {}

func (fs *FileServer) Start() {
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}
	
	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		
		go fs.ReadLoop(connection)
	}
}

func (fs *FileServer) ReadLoop(conn net.Conn) {
	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		file := buf[:n]
		fmt.Println(file)
		fmt.Printf("Recieved %d bytes over network \n", n)
	}
}