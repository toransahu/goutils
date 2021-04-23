/*
stack.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// Package "adt" is holder of "Abstract Data Types"
package adt

import "fmt"

// Stack implemented using slice of interface (should hold any data type)
type Stack []interface{}

// NewStack creates & returns an empty Stack
func NewStack() Stack {
	return []interface{}{}
}

// isNil (private) errors if the Stack is nil
func (s *Stack) isNil() error {
	if *s == nil {
		return fmt.Errorf("Stack %v is nil", s)
	}
	return nil
}

// IsEmpty tells whether the Stack is empty or not
func (s *Stack) IsEmpty() (bool, error) {
	if err := s.isNil(); err != nil {
		return false, err
	}
	return len(*s) == 0, nil
}

// Push inserts an item at the top of the Stack
func (s *Stack) Push(item interface{}) error {
	if err := s.isNil(); err != nil {
		return err
	}

	*s = append(*s, item)
	return nil
}

// Pop deletes & returns the item at the top of the Stack
func (s *Stack) Pop() (interface{}, error) {
	isEmpty, err := s.IsEmpty()
	if err != nil {
		return nil, err
	}
	if isEmpty {
		return nil, fmt.Errorf("Stack %v is empty", s)
	}

	top := (*s)[len(*s)-1] // get last element
	*s = (*s)[:len(*s)-1]  // remove last element from original slice
	return top, nil
}

// Top returns the item at the top of the Stack
func (s *Stack) Top() (interface{}, error) {
	isEmpty, err := s.IsEmpty()
	if err != nil {
		return nil, err
	}
	if isEmpty {
		return nil, fmt.Errorf("Stack %v is empty", s)
	}
	return (*s)[len(*s)-1], nil
}