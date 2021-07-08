/*
heap.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.

A good discussion & exercise on Inheretance Vs Composition in Golang.
https://github.com/golang/go/issues/22013
*/

// Package heap implements Heap using Array representation of a binary tree
package heap

import "math"

// Interface describes the requirements for a type using the functions/routines in this package.
type Interface interface {
	LessThan(int, int) bool
	Len() int
	Push(interface{})
	Pop() interface{}
	Swap(int, int)
	ItemAt(int) interface{}
	Set(int, interface{})
}

// Build (aka heapify) arranges the given unordered iterable of items such that it follows the property of min-heap.
// Time Complexity:
//	- Worst: O(n log n)
//	- Amortized: O(n)
// A good read on: a) why heapify() is O(n); b) why use percolate_down() instead percolate_up() in heapify()
// Ref: https://stackoverflow.com/questions/9755721/how-can-building-a-heap-be-on-time-complexity
func Build(arr Interface) {
	heapify(arr)
}

// Insert inserts a node in the heap in a correct order
// Time Complexity: O(log n) or better say O(no. of height/edge the node has to sift up)
func Insert(arr Interface, node interface{}) {
	arr.Push(node)
	lastNodePos := arr.Len() - 1
	percolateUp(arr, lastNodePos)
}

// Top (aka Min) returns top node of the heap
// Time Complexity: O(1)
func Top(arr Interface) interface{} {
	return arr.ItemAt(0)
}

// DeleteTop (aka ExtractMin) deletes the top node of the heap and returns the same
// Time Complexity: O(log n) or better say O(no. of height/edge the node has to sift down)
func DeleteTop(arr Interface) interface{} {
	// get the Top (Min) node
	top := Top(arr)

	// replace the top node with the last leaf
	poppedLastLeaf := arr.Pop()
	arr.Set(0, poppedLastLeaf)

	// percolateDown the new root node
	percolateDown(arr, 0)

	return top
}

// Replace deletes the top node of the heap and fills that with the given node.
// This is not same as DeleteTop() followed by Insert(); Replace() is more efficient as it avoid one round of percolateUp()
// Time Complexity: O(log n) or better say O(no. of height/edge the node has to sift down)
func Replace(arr Interface, node interface{}) interface{} {
	top := arr.ItemAt(0)
	arr.Set(0, node)
	percolateDown(arr, 0)
	return top
}

/*
 INTERNALS
*/

// heapify arranges the given unordered iterable of items such that it follows the property of min-heap.
// Time Complexity:
//	- Worst: O(n log n)
//	- Amortized: O(n)
// A good read on: a) why heapify() is O(n); b) why use percolate_down() instead percolate_up() in heapify()
// Ref: https://stackoverflow.com/questions/9755721/how-can-building-a-heap-be-on-time-complexity
func heapify(arr Interface) {
	// iterate the nodes from last (from right to left in the heap array)
	for idx := arr.Len() - 1; idx >= 0; idx-- {
		// percolateDown the node
		percolateDown(arr, idx)
	}
}

// percolateDown (aka siftDown) move a node down in the tree, similar to siftUp; used to restore heap condition after deletion or replacement.
// Time Complexity: O(log n) or better say O(no. of height/edge the node has to sift down)
func percolateDown(arr Interface, i int) {
	// get the size of the heap
	size := arr.Len()

	// if the given node is the leaf node or some non-existing node
	if i >= size/2 {
		// then stop the process
		return
	}

	// suppose minimun Node index is the given node's
	minimunNodePos := i
	// get the left child index
	leftChildPos := heapLeftChildPos(i)
	// get the right child index
	rightChildPos := heapRightChildPos(i)

	// if left child exists and if left child is less than the node at minimunNodePos
	if leftChildPos < size && arr.LessThan(leftChildPos, minimunNodePos) {
		minimunNodePos = leftChildPos
	}
	// if right child exists and if right child is less than the node at minimunNodePos
	if rightChildPos < size && arr.LessThan(rightChildPos, minimunNodePos) {
		minimunNodePos = rightChildPos
	}

	// if the given node is not the minimun node
	if i != minimunNodePos {
		// then swap the node with the minimunNodePos child
		arr.Swap(minimunNodePos, i)
		// and percolateDown the node at index minimunNodePos
		percolateDown(arr, minimunNodePos)
	}
}

// percolateUp (aka siftUp) move a node up in the tree, as long as needed; used to restore heap condition after insertion.
// Approach: Recursive (could be implemented iteratively as well)
// Time Complexity: O(log n) or better say O(no. of height/edge the node has to sift up)
func percolateUp(arr Interface, i int) {
	// if the given node is the root node or some non-existing node
	if i <= 0 { // TODO check the codition
		// exist the process
		return
	}

	// get the index of the parent of the given node
	parentPos := heapParentPos(i)

	if arr.LessThan(i, parentPos) {
		arr.Swap(i, parentPos)
		percolateUp(arr, parentPos)
	}
}

/*
 HELPERS
*/
func heapLeftChildPos(i int) int  { return 2*i + 1 }
func heapRightChildPos(i int) int { return 2*i + 2 }
func heapParentPos(i int) int     { return int(math.Floor(float64(i-1) / 2)) }
