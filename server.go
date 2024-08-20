package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
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

func (fs *FileServer) old_readLoop(conn net.Conn) {
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

func (fs *FileServer) ReadLoop(conn net.Conn) {
	buf := new(bytes.Buffer)
	for {
		var size int64
		err := binary.Read(conn, binary.LittleEndian, &size)
		if err != nil {
			return
		}
		n, err := io.CopyN(buf, conn, size)
		if err != nil {
			log.Fatal(err)
		}
		// panic("should panic")
		fmt.Println(buf.Bytes())
		fmt.Printf("Recieved %d bytes over network \n", n)
	}
}