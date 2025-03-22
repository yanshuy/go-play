package main

import "log"

func main() {
	server := NewServer()

	go server.handleCmds()

	log.Fatal(server.Start())
}
