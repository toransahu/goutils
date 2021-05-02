/*
linkedlist.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// This file implements Linkedlist

package adt

import (
	myerr "github.com/toransahu/goutils/errors"
)

var ERR_SLL_IS_EMPTY myerr.UserDefinedError = "single linkedlist is empty"
var ERR_SLLNODE_IS_NIL myerr.UserDefinedError = "sll node is nil"
var ERR_SLLNODE_DOES_NOT_EXISTS myerr.UserDefinedError = "sll node does not exists"

// SLLNode represents a Node of a Singly Linkedlist
type SLLNode struct {
	Data interface{} // Data/Value of the node
	Next *SLLNode    // pointer to the next node
}

// NewSLLNode creates and returns a SLLNode
func NewSLLNode(data interface{}) *SLLNode {
	return &SLLNode{Data: data}
}

// SLinkedlist represents a Singly Linkedlist
type SLinkedlist struct {
	Head *SLLNode // head of the Singly Linkedlist
}

// NewSLinkedList creates & returns a new Singly Linkedlist
func NewSLinkedList() SLinkedlist {
	return SLinkedlist{} // Head is nil
}

// IsEmpty tells whether a Singly Linkedlist is empty or not
func (l *SLinkedlist) IsEmpty() bool {
	return l.Head == nil
}

// Append adds a node at the end of the Singly Linkedlist
func (l *SLinkedlist) Append(data interface{}) *SLLNode {
	node := NewSLLNode(data)
	if l.IsEmpty() {
		l.Head = node
		return node
	}
	curr := l.Head
	for {
		if curr.Next == nil {
			curr.Next = node
			return curr.Next
		}
		curr = curr.Next
	}
}

// InsertAfter adds the given node after the given node of the Singly Linkedlist
func (l *SLinkedlist) InsertAfter(after *SLLNode, data interface{}) *SLLNode {
	node := NewSLLNode(data)
	tmp := after.Next
	after.Next = node
	node.Next = tmp
	return node
}

// InsertBefore adds the given node before the given node of the Singly Linkedlist
func (l *SLinkedlist) InsertBefore(before *SLLNode, data interface{}) *SLLNode {
	node := NewSLLNode(data)
	node.Next = before
	before = node
	return node
}

// InsertBeginning adds the given node the beginning of the Singly Linkedlist
func (l *SLinkedlist) InsertBeginning(data interface{}) *SLLNode {
	node := NewSLLNode(data)
	node.Next = l.Head
	l.Head = node
	return l.Head
}

// Remove deletes the given node from the Singly Linkedlist
func (l *SLinkedlist) Remove(node *SLLNode) error {
	if node == nil {
		return ERR_SLLNODE_IS_NIL
	}
	if l.IsEmpty() {
		return ERR_SLL_IS_EMPTY
	}

	if l.Head == node {
		l.Head = l.Head.Next
		return nil
	}

	curr := l.Head
	for {
		if curr.Next == nil {
			break
		}
		if curr.Next == node {
			curr = curr.Next
			return nil
		}
		curr = curr.Next
	}

	return ERR_SLLNODE_DOES_NOT_EXISTS
}

func (l *SLinkedlist) Iterate() <-chan *SLLNode {
	ch := make(chan *SLLNode)
	if l.IsEmpty() {
		defer close(ch)
		return ch
	}

	go func() {
		defer close(ch)
		curr := l.Head
		for {
			if curr.Next == nil {
				ch <- curr
				break
			}
			ch <- curr
			curr = curr.Next
		}
	}()
	return ch
}

// TODO: implement Doubly Linkedlist (List)
// TODO: implement Circular Linkedlist (Ring)
