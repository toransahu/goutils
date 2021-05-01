/*
linkedlist.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// This file implements Linkedlist

package adt

import myerr "github.com/toransahu/goutils/errors"

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
	for {
		next := l.Head.Next
		if next == nil {
			next = node
			return node
		}
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
	return node
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

	next := l.Head.Next
	for {
		if next == nil {
			break
		}
		if next == node {
			next = next.Next
			return nil
		}
		next = next.Next
	}

	return ERR_SLLNODE_DOES_NOT_EXISTS
}

// TODO: implement Doubly Linkedlist (List)
// TODO: implement Circular Linkedlist (Ring)
