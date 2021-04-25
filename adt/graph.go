/*
graph.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// This file implements Graphs
// Roadmap: provide interface; option to create new graph based on implementation type (e.g. AdjacencyMatrix, AdjacencyList etc.); Option to init the graph; create diff. types of graph (e.g. Directed, Undirected etc.) - where attributes/operations are diff.

package adt

// Graph denotes a Graph data structure
type Graph struct {
	Vertices      int
	AdjacencyList [][]int
}

// NewGraph creates & returns a Directed Graph implemented using Adjacency List
func NewGraph(numOfVertices int) *Graph {
	// create graph struct
	graph := &Graph{Vertices: numOfVertices, AdjacencyList: make([][]int, numOfVertices)}

	// init graph.AdjacencyMatrix with Zeros of [numOfVertices X numOfVertices] matrix
	for idx := range graph.AdjacencyList {
		// graph.AdjacencyList[idx] = make([]int, numOfVertices)
		graph.AdjacencyList[idx] = []int{}
	}
	return graph
}

// AddEdge inserts edge to the directed graph
func (g *Graph) AddEdge(u int, v int) {
	g.AdjacencyList[u] = append(g.AdjacencyList[u], v)
}

// DFS traverse the graph in Depth First Order and returns the vertices in the order
func (g *Graph) DFS() []int {
	// to store the ordered vertices
	result := []int{}
	// a memory map to flag the visited vertices
	visited := map[int]bool{}
	// arbitrarily choose first vertex in the AdjacencyList to start DFS
	vertex := 0 // Optional: for readability
	// run the DFS algo for `vertex` using `visited` memory
	// and append the result to `result`
	g.dfs(vertex, &visited, &result) // Optional: for readability

	// as this is a directed graph (and may be disconnected as well)
	// there could be possibilities that a few vertices remain unreachable
	// so in such case, iterate over all the vertices
	for vertex := range g.AdjacencyList {
		// and if they are not yet visited
		if !visited[vertex] {
			// run the DFS algo for the `vertex`
			g.dfs(vertex, &visited, &result)
		}
	}

	return result
}

// dfs (private func) does the basic Depth First Traversal of the graph
func (g *Graph) dfs(vertex int, visited *map[int]bool, result *[]int) {
	// create stack to help in backtracking whenever required
	// however, we can use recursion as well - it is same as using an explicit stack
	// as the recursive function calls will be stacked automatically

	// if the given vertex has already been visited
	if (*visited)[vertex] {
		return
	}
	// mark the given vertex as visited
	(*visited)[vertex] = true
	// append the given vertex to result list
	*result = append(*result, vertex)

	// iterate through all the adjacent vertices of the given vertex
	for _, u := range g.AdjacencyList[vertex] {
		// and run DFS for the adjacent vertex
		g.dfs(u, visited, result)
	}
}

// TopoSort sorts the directed acyclic graph (DAG) into Topological order
// A topological ordering is possible if and only if the graph has no directed cycles
func (g *Graph) TopoSort() []int {
	// to store the topological ordered vertices
	result := []int{}
	// stack to store the topological ordered vertices (in reverse order)
	stack := NewStack()
	// a memory map to flag the visited vertices
	visited := map[int]bool{}

	// iterate through all the vertices in the given graph
	// because the graph could be disconnected
	// or there could be some unreachable vertices if we start DFS/TopoSort from
	// a random vertex
	for vertex := range g.AdjacencyList {
		// and if they are not yet visited
		if !visited[vertex] {
			// and run the TopoSort (a tweaked DFS) for the given vertex
			g.topoSort(vertex, &visited, &stack)
		}
	}

	// if we're here, means all the vertices have been visited

	// while stack is not empty
	for {
		isEmpty, err := stack.IsEmpty()
		if err != nil {
			panic(err)
		}
		if isEmpty {
			break
		}
		// pop the vertex from the stack
		popped, err := stack.Pop()
		if err != nil {
			panic(err)
		}
		// store the popped vertex
		result = append(result, popped.(int))
	}

	return result
}

func (g *Graph) topoSort(vertex int, visited *map[int]bool, stack *Stack) {
	// if the given vertext has already been visited
	if (*visited)[vertex] {
		return
	}
	// mark the given vertex as visited
	(*visited)[vertex] = true

	// iterate through all the adjacent vertices of the given vertex
	for _, u := range g.AdjacencyList[vertex] {

		// and run Topo Sort (a little tweaked DFS - with a Stack) for the adjacent vertex
		g.topoSort(u, visited, stack)

	}

	// push the vertex into the stack
	stack.Push(vertex)
}
