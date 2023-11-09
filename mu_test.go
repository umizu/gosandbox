package main

import "testing"

func TestSetState(t *testing.T) {
	s := &State{}
	for i := 0; i < 10; i++ {
		go func(i int) {
			s.setState(i)
		}(i)
	}
}
