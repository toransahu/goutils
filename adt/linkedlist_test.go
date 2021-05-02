/*
linkedlist_test.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// Package adt
package adt

import (
	"reflect"
	"testing"
)

func TestSLinkedList_NewSLinkedList(t *testing.T) {
	l := NewSLinkedList()
	if l.Head != nil {
		t.Errorf("want: true; got: false")
	}
}

func TestSLinkedList_IsEmpty(t *testing.T) {
	l := NewSLinkedList()
	if !l.IsEmpty() {
		t.Errorf("want: true; got: false")
	}
}

func TestSLinkedList_Append(t *testing.T) {
	l := NewSLinkedList()
	l.Append(1)
	if l.IsEmpty() {
		t.Errorf("want: true; got: false")
	}
	if got := l.Head.Data; got != 1 {
		t.Errorf("want: 1; got: %v", got)
	}
	l.Append(2)
	if got := l.Head.Next.Data; got != 2 {
		t.Errorf("want: 2; got: %v", got)
	}
}

func TestSLinkedList_Iterate(t *testing.T) {
	l := NewSLinkedList()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)

	got := [][]int{}
	for i := range l.Iterate() {
		currData := i.Data.(int)
		next := i.Next
		var nextData int = -1
		if next != nil {
			nextData = i.Next.Data.(int)
		}
		got = append(got, []int{currData, nextData})
	}
	wanted := [][]int{
		{1, 2}, {2, 3}, {3, 4}, {4, -1},
	}
	if !reflect.DeepEqual(got, wanted) {
		t.Error()
	}
}

func TestSLinkedList_InsertAfter(t *testing.T) {
}

func TestSLinkedList_InsertBefore(t *testing.T) {
}

func TestSLinkedList_InsertBeginning(t *testing.T) {
}

func TestSLinkedList_Remove(t *testing.T) {
	// given - an empty linkedlist
	l := NewSLinkedList()
	// when - a nil node is given to be removed
	err := l.Remove(nil)
	// then
	if err != ERR_SLLNODE_IS_NIL {
		t.Errorf("want: %v; got: %v", ERR_SLLNODE_IS_NIL, err)
	}

	// given - an empty linkedlist
	// when - some node is given to be removed
	node := NewSLLNode(0)
	err = l.Remove(node)
	// then
	if err != ERR_SLL_IS_EMPTY {
		t.Errorf("want: %v; got: %v", ERR_SLL_IS_EMPTY, err)
	}

	// given - the linkedlist is not empty
	l.Append(1)
	// when - a nil node is given to be removed
	err = l.Remove(nil)
	// then
	if err != ERR_SLLNODE_IS_NIL {
		t.Errorf("want: %v; got: %v", ERR_SLLNODE_IS_NIL, err)
	}

	// given - the linkedlist is not empty
	// when - an unknown node is given to be removed
	err = l.Remove(node)
	// then
	if err != ERR_SLLNODE_DOES_NOT_EXISTS {
		t.Errorf("want: %v; got: %v", ERR_SLLNODE_DOES_NOT_EXISTS, err)
	}

	// given - the linkedlist is not empty
	// when - a valid node is given to be removed
	node = l.Head
	err = l.Remove(node)
	// then
	if err != nil {
		t.Errorf("want: %v; got: %v", nil, err)
	}
}
