package stack

import "errors"

var ErrEmpty = errors.New("Stack is empty")

type Stack struct {
	stack []int
}

func New() *Stack {
	return &Stack{make([]int, 0)}
}

func (s *Stack) Push(val int) {
	s.stack = append(s.stack, val)
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, ErrEmpty
	}
	temp := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return temp, nil
}

func (s *Stack) IsEmpty() bool {
	if len(s.stack) > 0 {
		return false
	}
	return true
}
