package main

import (
	"fmt"
	"log"
	"time"
)

type Server struct {
	msgch  chan string
	quitch chan struct{}
}

func (s *Server) Run() {
	fmt.Println()
	for {
		select {
		case msg := <-s.msgch:
			fmt.Println(msg)
		case <-s.quitch:
			log.Fatal("quitch received")
		}
	}
}

func RunSelectFile() {
	s := &Server{
		msgch:  make(chan string),
		quitch: make(chan struct{}),
	}
	go s.Run()
	s.msgch <- "hello"
	s.quitch <- struct{}{}
	time.Sleep(1 * time.Second)
}
