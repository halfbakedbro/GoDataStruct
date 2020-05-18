package stack

import "errors"

type Stack struct {
	top int
	data [] interface {}
}

func (s *Stack)Push(value interface {}) {
	s.data = append(s.data, value)
	s.top++
}

func (s *Stack)Pop() (value interface{}, err error) {
	if s.top <= 0 {
		return (nil, errors.New("Stack is empty"))
	}

	value = s.data[s.top-1]
	s.data = s.data[:s.top -1] 
	s.top--
	return value, nil
}

func (s Stack)Peek() (value interface {}, err error) {

	if s.top <= 0 {
		return (nil, errors.New("stack is empty"))
	}

	return s.data[s.top-1], nil
}

func (s Stack)Length() int {
	return s.top
}