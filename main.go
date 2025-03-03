package main

import (
	"log"
	"net"
)

type Server struct{
	listenAddr string
	ln net.Listener
}

func NewServer(addr string) *Server{
	return &Server{
		listenAddr: addr,
	}
}

func (s *Server) Start() error{
	ln, err := net.Dial("tcp", s.listenAddr)
	if err != nil {
		log.Fatal("Could not Start the server")
	}
	ln.Close()
}

func main(){

}