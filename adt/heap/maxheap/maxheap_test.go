/*
maxheap_test.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// Package maxheap
package maxheap

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/toransahu/goutils/adt/heap"
)

func TestHeapHelper(t *testing.T) {

	assert.Equal(t, 1, heapLeftChildPos(0))
	assert.Equal(t, 3, heapLeftChildPos(1))
	assert.Equal(t, 5, heapLeftChildPos(2))
	assert.Equal(t, 2, heapRightChildPos(0))
	assert.Equal(t, 4, heapRightChildPos(1))
	assert.Equal(t, 6, heapRightChildPos(2))
	assert.Equal(t, 0, heapParentPos(1))
	assert.Equal(t, 0, heapParentPos(2))
	assert.Equal(t, 1, heapParentPos(3))
	assert.Equal(t, 1, heapParentPos(4))
	assert.Equal(t, 2, heapParentPos(5))
	assert.Equal(t, 2, heapParentPos(6))
}

func TestHeapInternal(t *testing.T) {
	arr := &heap.IntArray{3, 0, 4, 2, 0, 1}

	heapify(arr)
	assert.Equal(t, 4, arr.ItemAt(0))
	assert.Equal(t, 2, arr.ItemAt(1))
	assert.Equal(t, 3, arr.ItemAt(2))
	assert.Equal(t, 0, arr.ItemAt(3))
	assert.Equal(t, 0, arr.ItemAt(4))
	assert.Equal(t, 1, arr.ItemAt(5))
}

func TestHeap(t *testing.T) {
	arr := &heap.IntArray{3, 0, 4, 2, 0, 1}

	Build(arr)
	assert.Equal(t, 4, arr.ItemAt(0))
	assert.Equal(t, 2, arr.ItemAt(1))
	assert.Equal(t, 3, arr.ItemAt(2))
	assert.Equal(t, 0, arr.ItemAt(3))
	assert.Equal(t, 0, arr.ItemAt(4))
	assert.Equal(t, 1, arr.ItemAt(5))

	top := Top(arr)
	assert.Equal(t, 4, top)

	top = DeleteTop(arr)
	assert.Equal(t, 4, top)
	assert.Equal(t, 5, arr.Len())

	Insert(arr, 5)
	assert.Equal(t, 5, Top(arr))

	top = Replace(arr, -1)
	assert.Equal(t, 5, top)
	assert.Equal(t, 3, Top(arr))
	assert.Equal(t, -1, arr.ItemAt(5))
}
