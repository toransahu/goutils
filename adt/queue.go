/*
queue.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// This file implements queues

package adt

import "github.com/toransahu/goutils/errors"

var ERR_QUEUE_IS_EMPTY errors.UserDefinedError = "queue is empty"

type Queue []interface{}

func NewQueue() *Queue {
	var q Queue
	q = make([]interface{}, 0)
	return &q
}

func (q *Queue) Enqueue(item interface{}) {
	*q = append(*q, item)
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return nil, ERR_QUEUE_IS_EMPTY
	}
	item := (*q)[0]
	*q = (*q)[1:]
	return item, nil
}
