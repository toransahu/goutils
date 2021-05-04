/*
queue_test.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

package adt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue_NewQueue(t *testing.T) {
	q := NewQueue()
	assert.NotNil(t, q)
	assert.Equal(t, 0, len((*q)))
}

func TestQueue_Enqueue(t *testing.T) {
	q := NewQueue()
	assert.Equal(t, 0, len((*q)))
	q.Enqueue(1)
	assert.Equal(t, 1, len((*q)))
	assert.Equal(t, 1, (*q)[0])
	q.Enqueue(2)
	assert.Equal(t, 2, len((*q)))
	assert.Equal(t, 1, (*q)[0])
	assert.Equal(t, 2, (*q)[1])
}

func TestQueue_IsEmpty(t *testing.T) {
	q := NewQueue()
	assert.NotNil(t, q)
	assert.Equal(t, 0, len((*q)))
	assert.True(t, q.IsEmpty())
	q.Enqueue(1)
	assert.False(t, q.IsEmpty())
}

func TestQueue_Dequeue(t *testing.T) {
	q := NewQueue()
	q.Enqueue(1)
	item, _ := q.Dequeue()
	assert.Equal(t, 1, item.(int))

	q.Enqueue(1)
	q.Enqueue(2)
	item, _ = q.Dequeue()
	assert.Equal(t, 1, item.(int))
	item, _ = q.Dequeue()
	assert.Equal(t, 2, item.(int))
	item, err := q.Dequeue()
	assert.Equal(t, ERR_QUEUE_IS_EMPTY, err)
	assert.Nil(t, item)
}
