package main

import "time"

func main() {
	go func() {
		time.Sleep(2 * time.Second)
		SendFile(4000)
	}()
	server := &FileServer{}
	server.Start()
}
