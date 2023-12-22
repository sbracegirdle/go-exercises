package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

type Server struct {
	host string
	port string
}

func NewServer(host string, port string) *Server {
	return &Server{host: host, port: port}
}

func (s *Server) Start() {
	ln, err := net.Listen("tcp", s.host+":"+s.port)
	if err != nil {
		log.Fatalf("Failed to start server: %s\n", err)
		return
	}
	defer ln.Close()

	log.Printf("Server started on %s:%s\n", s.host, s.port)

	// Loop forever and accept incoming connections, handle them in a goroutine
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %s\n", err)
			continue
		}
		go s.handleRequest(conn)
	}
}

func (s *Server) handleRequest(conn net.Conn) {
	defer conn.Close()

	request, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Printf("Failed to read request: %s\n", err)
		return
	}

	fmt.Println("Received request: ", strings.TrimSpace(request))

	response := "HTTP/1.1 200 OK\r\n\r\nHello, World!"
	conn.Write([]byte(response))
}

func main() {
	server := NewServer("localhost", "1234")
	server.Start()
}
