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
}

func TestGraph_IsCyclic(t *testing.T) {
	g := NewGraph(2)
	g.AddEdge(0, 1)
	g.AddEdge(1, 0)
	result := g.IsCyclic()

	if !result {
		t.Errorf("for given %v, wanted: %v, got: %v", g, true, result)
	}
}
