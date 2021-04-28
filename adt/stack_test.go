//
// stack_test.go
// Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>
//
// Distributed under terms of the MIT license.
//

package adt

import (
	"testing"
)

func TestStack_isNil(t *testing.T) {
	var s Stack
	err := s.isNil()
	if err == nil {
		t.Fail()
	}

	s = NewStack()
	err = s.isNil()
	if err != nil {
		t.Fail()
	}
}

func TestStack_IsEmpty_Push_Pop_Top(t *testing.T) {
	s := NewStack()

	// test IsEmpty()
	isEmpty, err := s.IsEmpty()
	if err != nil {
		t.Fail()
	}
	if !isEmpty {
		t.Errorf("wanted: true; got: %v", isEmpty)
	}

	// test Push() (in following lines of code)
	s.Push(10)
	s.Push(20)

	// test IsEmpty()
	isEmpty, err = s.IsEmpty()
	if err != nil {
		t.Fail()
	}
	if isEmpty {
		t.Errorf("wanted: false; got: %v", isEmpty)
	}

	// test Top()
	top, err := s.Top()
	if top != 20 {
		t.Errorf("wanted: 20; got: %v", top)
	}

	// test Pop()
	popped, err := s.Pop()
	if popped != 20 {
		t.Errorf("wanted: 20; got: %v", popped)
	}
	popped, err = s.Pop()
	if popped != 10 {
		t.Errorf("wanted: 10; got: %v", popped)
	}
	popped, err = s.Pop()
	if err == nil {
		t.Errorf("wanted: %v; got: %v", ERR_STACK_IS_EMPTY, err)
	}
	popped, err = s.Pop()
	if err == nil {
		t.Errorf("wanted: %v; got: %v", ERR_STACK_IS_EMPTY, err)
	}

	// test Top()
	top, err = s.Top()
	if err == nil {
		t.Errorf("wanted: %v; got: %v", ERR_STACK_IS_EMPTY, err)
	}
}
