package simplestack

import (
	"container/list"
	"fmt"
)

type Stack struct {
	stack *list.List
}

type any interface{}

func NewStack() *Stack {
	return &Stack{stack: list.New()}
}

func (s *Stack) Empty() bool {
	return s.stack.Len() == 0
}

func (s *Stack) Size() int {
	return s.stack.Len()
}

func (s *Stack) Push(value any) {
	s.stack.PushFront(value)
}

func (s *Stack) Pop() {
	if s.Size() > 0 {
		item := s.stack.Front()
		s.stack.Remove(item)
	}
}

func (s *Stack) Peek() (any, error) {
	if s.Size() > 0 {
		if value, ok := s.stack.Front().Value.(any); ok {
			return value, nil
		}
		return nil, fmt.Errorf("Error")
	}
	return nil, fmt.Errorf("Error")
}
