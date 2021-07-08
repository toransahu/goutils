/*
maxheap.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// Package maxheap
package maxheap

import "math"

// Interface describes the requirements for a type using the functions/routines in this package.
type Interface interface {
	GreaterThan(int, int) bool
	Len() int
	Push(interface{})
	Pop() interface{}
	Swap(int, int)
	ItemAt(int) interface{}
	Set(int, interface{})
}

// Build (aka heapify) arranges the given unordered iterable of items such that it follows the property of max-heap.
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

// Top (aka Max) returns top node of the heap
// Time Complexity: O(1)
func Top(arr Interface) interface{} {
	return arr.ItemAt(0)
}

// DeleteTop (aka ExtractMax) deletes the top node of the heap and returns the same
// Time Complexity: O(log n) or better say O(no. of height/edge the node has to sift down)
func DeleteTop(arr Interface) interface{} {
	// get the Top (Max) node
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

// heapify arranges the given unordered iterable of items such that it follows the property of max-heap.
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

	// suppose maximum Node index is the given node's
	maximumNodePos := i
	// get the left child index
	leftChildPos := heapLeftChildPos(i)
	// get the right child index
	rightChildPos := heapRightChildPos(i)

	// if left child exists and if left child is less than the node at maximumNodePos
	if leftChildPos < size && arr.GreaterThan(leftChildPos, maximumNodePos) {
		maximumNodePos = leftChildPos
	}
	// if right child exists and if right child is less than the node at maximumNodePos
	if rightChildPos < size && arr.GreaterThan(rightChildPos, maximumNodePos) {
		maximumNodePos = rightChildPos
	}

	// if the given node is not the maximum node
	if i != maximumNodePos {
		// then swap the node with the maximumNodePos child
		arr.Swap(maximumNodePos, i)
		// and percolateDown the node at index maximumNodePos
		percolateDown(arr, maximumNodePos)
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

	if arr.GreaterThan(i, parentPos) {
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
