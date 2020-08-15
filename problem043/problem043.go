package problem043

import (
	"errors"
	"github.com/thoas/go-funk"
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

type stack struct {
	values   []int
	maxValue int
}

func newStack() stack {
	return stack{}
}

func (s *stack) push(value int) {
	s.values = append(s.values, value)
	s.updateMaxValue(&value)
}

func (s *stack) isEmpty() bool {
	return len(s.values) == 0
}

func (s *stack) updateMaxValue(value *int) {
	// REVIEW: こんなことしたくない
	if !s.isEmpty() {
		s.maxValue = funk.MaxInt(s.values).(int)
	}

	if value != nil && *value > s.maxValue {
		s.maxValue = *value
	}
}

func (s *stack) pop() (int, error) {
	if s.isEmpty() {
		return 0, errors.New("stack is empty")
	}

	pop := s.values[len(s.values) - 1]
	s.values = s.values[:len(s.values) - 1]
	s.updateMaxValue(nil)

	return pop, nil
}

func (s *stack) max() (int, error) {
	if s.isEmpty() {
		return 0, errors.New("stack is empty")
	}

	return s.maxValue, nil
}
