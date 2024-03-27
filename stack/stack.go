package stack

import "fmt"

type Stack struct {
	elements []int
}

func NewStack() *Stack {
	return &Stack{elements: []int{}}
}

func (s *Stack) Push(element int) {
	s.elements = append(s.elements, element)
}

func (s *Stack) Pop() (int, error) {
	if len(s.elements) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}

	lastIndex := len(s.elements) - 1
	topElement := s.elements[lastIndex]
	s.elements = s.elements[:lastIndex]
	return topElement, nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack) Top() (int, error) {
	if len(s.elements) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}

	return s.elements[len(s.elements)-1], nil
}
