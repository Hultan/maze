package maze_gen

import (
	"sync"
)

type Stack struct {
	lock sync.Mutex // you don't have to do this if you don't want thread safety
	s    []Location
}

func NewStack() *Stack {
	return &Stack{sync.Mutex{}, make([]Location, 0)}
}

func (s *Stack) Length() int {
	return len(s.s)
}

func (s *Stack) Push(l Location) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.s = append(s.s, l)
}

func (s *Stack) Pop() Location {
	s.lock.Lock()
	defer s.lock.Unlock()

	l := len(s.s)
	if l == 0 {
		return Location{}
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res
}
