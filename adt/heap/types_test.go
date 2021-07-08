/*
types_test.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// Package heap
package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntArray(t *testing.T) {
	arr := IntArray{1, 2, 3, 4}
	assert.True(t, arr.LessThan(0, 1))
	assert.Equal(t, len(arr), arr.Len())
	assert.Equal(t, arr[0], arr.ItemAt(0))

	popped := arr.Pop()
	assert.Equal(t, popped, 4)
	assert.Equal(t, 3, arr.Len())

	arr.Swap(0, 1)
	assert.Equal(t, 2, arr.ItemAt(0))
	assert.Equal(t, 1, arr.ItemAt(1))

	arr.Set(0, 5)
	assert.Equal(t, 5, arr.ItemAt(0))

	arr.Push(5)
	assert.Equal(t, 5, arr.ItemAt(3))
	assert.Equal(t, 4, arr.Len())
}
