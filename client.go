package main

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

func old_sendFile(size int) error {
	file := make([]byte, size)
	_, err := io.ReadFull(rand.Reader, file)
	if err != nil {
		return err
	}
	
	conn, err := net.Dial("tcp", ":3000")
	if err != nil {
		return err
	}
	n, err := conn.Write(file)
	if err != nil {
		return err
	}
	fmt.Printf("written %d byte over network \n", n)
	return nil
}


func SendFile(size int) error {
	file := make([]byte, size)
	_, err := io.ReadFull(rand.Reader, file)
	if err != nil {
		return err
	}
	
	conn, err := net.Dial("tcp", ":3000")
	if err != nil {
		return err
	}
	
	binary.Write(conn, binary.LittleEndian, int64(size))
	n, err := io.CopyN(conn, bytes.NewReader(file), int64(size))
	if err != nil {
		return err
	}
	fmt.Printf("written %d byte over network \n", n)
	return nil
}