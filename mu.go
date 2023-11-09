package main

import "sync"

type State struct {
	mu    sync.Mutex
	count int
}

func (s *State) setState(i int) {
	s.mu.Lock()
	s.count = i
	s.mu.Unlock()
}
