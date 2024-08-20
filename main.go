package main

import "time"

func main() {
	go func() {
		time.Sleep(2 * time.Second)
		SendFile(20000000) // number of bytes to be streamed over network
	}()
	server := &FileServer{}
	server.Start()
}
