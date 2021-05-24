/*
graph_test.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

package adt

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraph_dfs(t *testing.T) {
	g := NewGraph(5)
	g.AddEdge(0, 3)
	g.AddEdge(3, 4)
	g.AddEdge(4, 1)
	g.AddEdge(2, 1)
	g.AddEdge(2, 0)
	result := []int{}
	visited := map[int]bool{}
	g.dfs(0, &visited, &result)
	wanted := []int{0, 3, 4, 1}

	if !reflect.DeepEqual(result, wanted) {
		t.Errorf("for given %v, wanted: %v, got: %v", g, wanted, result)
	}
}

func TestGraph_DFS(t *testing.T) {
	g := NewGraph(5)
	g.AddEdge(0, 3)
	g.AddEdge(3, 4)
	g.AddEdge(4, 1)
	g.AddEdge(2, 1)
	g.AddEdge(2, 0)
	result := g.DFS()
	wanted := []int{0, 3, 4, 1, 2}

	if !reflect.DeepEqual(result, wanted) {
		t.Errorf("for given %v, wanted: %v, got: %v", g, wanted, result)
	}
}

func TestGraph_TopoSort(t *testing.T) {
	g := NewGraph(5)
	g.AddEdge(0, 3)
	g.AddEdge(3, 4)
	g.AddEdge(4, 1)
	g.AddEdge(2, 1)
	g.AddEdge(2, 0)
	result, hasCycle := g.TopoSort()
	wanted := []int{2, 0, 3, 4, 1}

	assert.False(t, hasCycle)

	if !reflect.DeepEqual(result, wanted) {
		t.Errorf("for given %v, wanted: %v, got: %v", g, wanted, result)
	}

	g = NewGraph(3)
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)
	result, hasCycle = g.TopoSort()
	assert.True(t, hasCycle)
	assert.Nil(t, result)

	g = NewGraph(6)
	g.AddEdge(5, 2)
	g.AddEdge(5, 0)
	g.AddEdge(2, 3)
	g.AddEdge(3, 1)
	g.AddEdge(4, 0)
	g.AddEdge(4, 1)
	result, hasCycle = g.TopoSort()
	assert.False(t, hasCycle)
	assert.Equal(t, []int{5, 4, 2, 3, 1, 0}, result)

}

func TestGraph_IsCyclic(t *testing.T) {
	/*

		1  -->  0
		^       |
		|_______|

	*/
	g1 := NewGraph(2)
	g1.AddEdge(0, 1)
	g1.AddEdge(1, 0)

	/*

			  	   5----->2
			      /        \
			     /          \
			    v            v
		        0            3
			    ^            ^
			     \          /
			      \        /
			       4----->1

	*/
	g2 := NewGraph(6)
	g2.AddEdge(5, 2)
	g2.AddEdge(5, 0)
	g2.AddEdge(2, 3)
	g2.AddEdge(3, 1)
	g2.AddEdge(4, 0)
	g2.AddEdge(4, 1)

	/*
		------> 0
		^       |
		|_______|

	*/
	g3 := NewGraph(1)
	g3.AddEdge(0, 0)

	/*
			    0
		  	   / ^
		      /   \
		     /     \
		    v       \
			1 -----> 2

	*/
	g4 := NewGraph(3)
	g4.AddEdge(0, 1)
	g4.AddEdge(1, 2)
	g4.AddEdge(2, 0)

	testcases := []struct {
		given *Graph
		want  bool
	}{
		{g1, true},
		{g2, false},
		{g3, true},
		{g4, true},
	}

	for _, tc := range testcases {

		got := tc.given.IsCyclic()
		if got != tc.want {
			t.Errorf("for given %v, wanted: %v, got: %v", tc.given, tc.want, got)
		}

		got = tc.given.IsCyclic_V2()
		if got != tc.want {
			t.Errorf("for given %v, wanted: %v, got: %v", tc.given, tc.want, got)
		}
	}
}
