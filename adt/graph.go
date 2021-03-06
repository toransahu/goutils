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
func (g *Graph) TopoSort() ([]int, bool) {
	// to store the topological ordered vertices
	result := []int{}
	// stack to store the topological ordered vertices (in reverse order)
	stack := NewStack()
	// a memory map to flag the visited vertices
	visited := map[int]bool{}
	// a memory map to detect cycle by flagging the recently visited vertices
	// i.e. in on going call stack, before any backtracking
	recentlyVisited := map[int]bool{}

	// iterate through all the vertices in the given graph
	// because the graph could be disconnected
	// or there could be some unreachable vertices when we start DFS/TopoSort from
	// a random vertex
	for vertex := range g.AdjacencyList {
		// and if the vertex is not yet visited
		if !visited[vertex] {
			// run the TopoSort (a tweaked DFS) for the given vertex
			hasCycle := g.topoSort(vertex, &visited, &stack, &recentlyVisited)
			if hasCycle {
				return nil, true
			}
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

	// return result and no cycle detected
	return result, false
}

func (g *Graph) topoSort(vertex int, visited *map[int]bool, stack *Stack, recentlyVisited *map[int]bool) bool {
	// if the given vertex has already been visited in the ongoing call stack
	// before backtracking
	if (*recentlyVisited)[vertex] {
		// then a cycle exists in the graph
		return true
	}

	// if the given vertex has already been visited
	if (*visited)[vertex] {
		return false
	}
	// mark the given vertex as visited
	(*visited)[vertex] = true

	// mark the given vertex as recentlyVisited
	(*recentlyVisited)[vertex] = true

	// iterate through all the adjacent vertices of the given vertex
	for _, u := range g.AdjacencyList[vertex] {

		// and run Topo Sort (a little tweaked DFS - with a Stack) for the adjacent vertex
		hasCycle := g.topoSort(u, visited, stack, recentlyVisited)
		if hasCycle {
			return true
		}
	}
	*recentlyVisited = map[int]bool{}

	// push the vertex into the stack
	stack.Push(vertex)
	return false
}

func (g *Graph) IsCyclic() bool {
	// a memory map to flag the visited vertices
	visited := map[int]bool{}
	// a map to track all the vertices we visited before a backtrack
	recentlyVisited := map[int]bool{}
	for vertex := range g.AdjacencyList {
		// as this is a directed graph (and may be disconnected as well)
		// there could be possibilities that a few vertices remain unreachable
		// so in such case, iterate over all the vertices
		isCyclic := g.isCyclic(vertex, &visited, &recentlyVisited)
		if isCyclic {
			return true
		}
	}
	return false
}

func (g *Graph) isCyclic(vertex int, visited *map[int]bool, recentlyVisited *map[int]bool) bool {
	// if the given vertex has already been visited in __current call stack__
	if (*recentlyVisited)[vertex] {
		// then there is a loop
		return true
	}

	// if the given vertex has already been visited
	if (*visited)[vertex] {
		return false
	}
	// mark the given vertex as visited
	(*visited)[vertex] = true
	(*recentlyVisited)[vertex] = true

	// iterate through all the adjacent vertices of the given vertex
	for _, u := range g.AdjacencyList[vertex] {
		// and run DFS for the adjacent vertex
		isCyclic := g.isCyclic(u, visited, recentlyVisited)
		if isCyclic {
			return true
		}
	}
	*recentlyVisited = map[int]bool{}
	return false
}

// IsCyclic_V2 detects cycle in a directed graph using BFS by maintaing 3 colors of each node
func (g *Graph) IsCyclic_V2() bool {
	// a memory map to flag the visited vertices
	visited := map[int]bool{}
	for vertex := range g.AdjacencyList {
		if visited[vertex] {
			continue
		}
		// as this is a directed graph (and may be disconnected as well)
		// there could be possibilities that a few vertices remain unreachable
		// so in such case, iterate over all the vertices
		isCyclic := g.isCyclic_V2(vertex, &visited)
		if isCyclic {
			return true
		}
	}
	return false
}

func (g *Graph) isCyclic_V2(vertex int, visited *map[int]bool) bool {
	// a memory map to flag the visited vertices
	colors := map[int]int{}

	q := NewQueue()
	q.Enqueue(vertex)

	colors[vertex] = 0
	(*visited)[vertex] = true

	for {
		if q.IsEmpty() {
			break
		}

		n, err := q.Dequeue()
		if err != nil {
			panic(err)
		}
		node := n.(int)
		colors[node] = -1
		(*visited)[vertex] = true

		for _, neighbor := range g.AdjacencyList[node] {
			if colors[neighbor] == -1 {
				return true
			}
			if (*visited)[neighbor] {
				continue
			}
			q.Enqueue(neighbor)
			colors[neighbor] = 0
		}
	}
	return false
}

// IsCyclic_V3 detects cycle in a directed graph using BFS by manupulating (reducing) in-degree of the node
// Idea: If there exists a cycle, then the vertices involved in the cycle would have in-degree greater than zero. So, if we remove all the vertices having in-degree == 0 in the graph and find that the graph is left with all the vertices having in-degree > 0; we can conclude the graph have a cycle.
func (g *Graph) IsCyclic_V3() bool {
	// to hold the visited vertices
	visited := map[int]bool{}

	// to hold the in-degrees of each vertex
	inDegreeMap := map[int]int{}

	// calculate in-degrees
	for _, neighbors := range g.AdjacencyList {
		for _, neighbor := range neighbors {
			inDegreeMap[neighbor]++
		}
	}

	// pre-check
	// if in-degree of all the vertices are > 0 then declare the graph cyclic
	for _, inD := range inDegreeMap {
		if inD != 0 {
			continue
		}
		return true
	}

	// number of vertices we successfully removed from the graph (whose in-degree is/become zero)
	removed := 0

	// as this is a directed graph (and may be disconnected as well)
	// there could be possibilities that a few vertices remain unreachable
	// so in such case, iterate over all the vertices
	for vertex := range g.AdjacencyList {
		// but run the check for only those vertices which are un-visited and have in-degree == 0
		if visited[vertex] || inDegreeMap[vertex] != 0 {
			continue
		}
		g.isCyclic_V3(vertex, visited, &removed, inDegreeMap)
		if removed == len(g.AdjacencyList) {
			return false
		}
	}
	return true
}

func (g *Graph) isCyclic_V3(vertex int, visited map[int]bool, removed *int, inDegreeMap map[int]int) {
	q := NewQueue()
	q.Enqueue(vertex)

	for {
		if q.IsEmpty() {
			break
		}

		n, err := q.Dequeue()
		if err != nil {
			panic(err)
		}
		node := n.(int)
		visited[node] = true
		// as we removed this vertex from the graph, increase the counter
		(*removed)++

		for _, neighbor := range g.AdjacencyList[node] {
			// as we removed the parent of this vertex, decrease its in-degree by 1
			inDegreeMap[neighbor]--
			inDegree := inDegreeMap[neighbor]
			// if thi vertex's in-degree became zero, now we need to do the whole procedure for this vertex as well
			if inDegree == 0 {
				// so enqueue it
				q.Enqueue(neighbor)
			}
		}
	}
}

/* TODO - Better design

type Graph interface{}
type GraphNode interface{}

func (g *Graph) AddEdge(u GraphNode, v GraphNode, weight int)

NewSimpleGraph(implType=it.AdjacencyList) == NewGraph(implType=it.AdjacencyList)
NewGraph(implType=it.AdjacencyList)

// DGNode
type DGNode struct {
}

// DGraph denotes a Directed Graph data structure
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
*/
