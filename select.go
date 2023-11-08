package main

import (
	"fmt"
	"log"
)

type Server struct {
	msgch  chan string
	quitch chan struct{}
}

func (s *Server) Run() {
	for {
		select {
		case <-s.msgch:
			fmt.Println("msg received")
		case <-s.quitch:
			log.Fatal("quitch received")
		}
	}
}

func main() {
	s := &Server{
		msgch: make(chan string),
		quitch: make(chan struct{}),
	}
	go s.Run()
	s.quitch <- struct{}{}
}
