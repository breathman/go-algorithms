package stack

import "errors"

type Stack struct {
	elements []int
}

func New() *Stack {
	return &Stack{}
}

func (s *Stack) Push(item int) {
	s.elements = append(s.elements, item)
}

func (s *Stack) Pop() (int, error) {
	if !s.IsEmpty() {
		res := s.elements[s.Size()-1]
		s.elements = s.elements[:s.Size()-1]
		return res, nil
	}
	return 0, errors.New("stack is empty")
}

func (s *Stack) Top() (int, error) {
	if !s.IsEmpty() {
		return s.elements[s.Size()-1], nil
	}
	return 0, errors.New("stack is empty")
}

func (s *Stack) Size() int {
	return len(s.elements)
}

func (s *Stack) IsEmpty() bool {
	if s.Size() == 0 {
		return true
	}
	return false
}
