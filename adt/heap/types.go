/*
types.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// Package heap
package heap

type IntArray []int

func (a IntArray) LessThan(i, j int) bool      { return a[i] < a[j] }
func (a IntArray) Len() int                    { return len(a) }
func (a IntArray) Swap(i, j int)               { a[i], a[j] = a[j], a[i] }
func (a IntArray) ItemAt(i int) interface{}    { return a[i] }
func (a IntArray) Set(i int, item interface{}) { a[i] = item.(int) }
func (a *IntArray) Push(item interface{})      { *a = append(*a, item.(int)) }
func (a *IntArray) Pop() interface{} {
	lastIndex := len(*a) - 1
	popped := (*a)[lastIndex]
	*a = (*a)[0:lastIndex]
	return popped
}
