package stack

import (
	"../linklist"
)

/**
stack implementation with List
**/

type Stack struct {
	list linklist.LinkedList
}

func initStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(value interface{}) error {
	return s.list.Prepend(value)
}

func (s *Stack) Pop() (interface{}, error) {
	return s.list.PopFront()
}

func (s Stack) Length() int {
	return s.list.Len
}
