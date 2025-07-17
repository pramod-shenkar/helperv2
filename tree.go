package main

import (
	"fmt"
)

// Stack represents a generic stack data structure.
type Stack[T any] struct {
	elements []T
}

// NewStack creates and returns a new empty Stack.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		elements: make([]T, 0),
	}
}

// Push adds an element to the top of the stack.
func (s *Stack[T]) Push(item T) {
	s.elements = append(s.elements, item)
}

// Pop removes and returns the element from the top of the stack.
// It returns an error if the stack is empty.
func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("stack is empty, cannot pop")
	}
	index := len(s.elements) - 1
	item := s.elements[index]
	s.elements = s.elements[:index]
	return item, nil
}

// IsEmpty checks if the stack is empty.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}
