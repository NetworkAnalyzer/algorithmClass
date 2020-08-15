package problem043

import (
	"errors"
)

// This problem was asked by Amazon.
//
// Implement a stack that has the following methods:
//
// push(val), which pushes an element onto the stack
// pop(), which pops off and returns the topmost element of the stack.
//        If there are no elements in the stack, then it should throw an error or return null.
// max(), which returns the maximum value in the stack currently.
//        If there are no elements in the stack, then it should throw an error or return null.
//
// Each method should run in constant time.

type Stack struct {
	values     []int
	maxIndices []int
}

func NewStack() Stack {
	return Stack{}
}

func (s *Stack) Push(value int) {
	s.values = append(s.values, value)
	s.updateMaxIndices(value)
}

func (s *Stack) Pop() (int, error) {
	if empty(s.values) {
		return 0, errors.New("stack is empty")
	}

	if lastIndex(s.values) == last(s.maxIndices) {
		s.maxIndices = removeLast(s.maxIndices)
	}

	pop := last(s.values)
	s.values = removeLast(s.values)

	return pop, nil
}

func (s *Stack) Max() (int, error) {
	if empty(s.values) {
		return 0, errors.New("stack is empty")
	}

	return s.getMaxValue(), nil
}

func (s *Stack) updateMaxIndices(value int) {
	if empty(s.maxIndices) || value > s.getMaxValue() {
		s.maxIndices = append(s.maxIndices, lastIndex(s.values))
	}
}

func (s *Stack) getMaxValue() int {
	return s.values[last(s.maxIndices)]
}

// utils
func empty(slice []int) bool {
	return len(slice) == 0
}

func last(slice []int) int {
	return slice[lastIndex(slice)]
}

func lastIndex(slice []int) int {
	return len(slice) - 1
}

func removeLast(slice []int) []int {
	return slice[:len(slice) - 1]
}
