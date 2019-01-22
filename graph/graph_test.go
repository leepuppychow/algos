package graph

import (
	"fmt"
	"testing"
)

var v0 = Vertex{Data: "0"}
var v1 = Vertex{Data: "1"}
var v2 = Vertex{Data: "2"}
var v3 = Vertex{Data: "3"}
var v4 = Vertex{Data: "4"}
var v5 = Vertex{Data: "5"}

var e0 = Edge{V1: v0, V2: v1, Distance: 10}
var e1 = Edge{V1: v0, V2: v3, Distance: 20}
var e2 = Edge{V1: v0, V2: v5, Distance: 5}
var e3 = Edge{V1: v1, V2: v2, Distance: 16}
var e4 = Edge{V1: v2, V2: v5, Distance: 31}
var e5 = Edge{V1: v3, V2: v4, Distance: 8}
var e6 = Edge{V1: v3, V2: v5, Distance: 7}
var e7 = Edge{V1: v4, V2: v5, Distance: 19}

func Setup() Graph {
	adjList := map[Vertex][]Vertex{
		v0: []Vertex{v1, v3, v5},
		v1: []Vertex{v0, v2},
		v2: []Vertex{v1, v5},
		v3: []Vertex{v0, v4, v5},
		v4: []Vertex{v3, v5},
		v5: []Vertex{v0, v2, v3, v4},
	}

	visited := map[Vertex]bool{
		v0: true,
		v5: true,
	}

	blankParents := map[Vertex]Vertex{}

	return Graph{
		AdjacencyList: adjList,
		EdgeList:      []Edge{e0, e1, e2, e3, e4, e5, e6, e7},
		Visited:       visited,
		VisitQueue:    []Vertex{v1, v2},
		Parents:       blankParents,
	}
}

func TestIncludes(t *testing.T) {
	graph1 := Setup()
	if !Includes(graph1.VisitQueue, v1) {
		t.Errorf("Includes test failed")
	}

	if Includes(graph1.VisitQueue, v5) {
		t.Errorf("Includes test failed")
	}
}

func TestClear(t *testing.T) {
	graph1 := Setup()
	graph1.Clear()

	if len(graph1.Visited) != 0 || len(graph1.VisitQueue) != 0 {
		t.Errorf("Clear Method failed")
	}
}

func TestBreadthFirstSearch(t *testing.T) {
	testGraph := Setup()
	testGraph.Clear()
	vNope := Vertex{Data: "Blah"}

	fmt.Println("***********BREADTH FIRST SEARCH RUNNING**************")
	if !testGraph.BreadthFirstSearch(v0, v3) {
		t.Errorf("BFS method failed")
	}

	testGraph.Clear()
	fmt.Println("***********BREADTH FIRST SEARCH RUNNING**************")
	if testGraph.BreadthFirstSearch(v0, vNope) {
		t.Errorf("BFS method failed")
	}
}

func TestDepthFirstSearch(t *testing.T) {
	testGraph := Setup()
	testGraph.Clear()
	vNope := Vertex{Data: "Blah"}

	fmt.Println("~~~~~~~~ DEPTH FIRST SEARCH RUNNING ~~~~~~~~~~~~~")
	if !testGraph.DepthFirstSearch(v0, v3) {
		t.Errorf("DFS method failed")
	}

	testGraph.Clear()
	fmt.Println("~~~~~~~~ DEPTH FIRST SEARCH RUNNING ~~~~~~~~~~~~~")
	if testGraph.DepthFirstSearch(v0, vNope) {
		t.Errorf("DFS method failed")
	}
}

func TestDepth(t *testing.T) {
	testGraph := Setup()
	vNope := Vertex{Data: "Blah"}

	testGraph.Clear()
	fmt.Println("***********BREADTH FIRST SEARCH RUNNING**************")
	testGraph.BreadthFirstSearch(v2, vNope)

	if testGraph.Depth(v2) != 0 {
		t.Errorf("Depth failed")
	}

	if testGraph.Depth(v5) != 1 {
		t.Errorf("Depth failed")
	}

	if testGraph.Depth(v0) != 2 {
		t.Errorf("Depth failed")
	}
}

func TestEdgeExists(t *testing.T) {
	g := Setup()

	yes, distance := g.EdgeExists(v0, v1)
	if !yes || distance != 10 {
		t.Errorf("EdgeExists failed")
	}
	// Undirected graph for now
	yes2, distance2 := g.EdgeExists(v1, v0)
	if !yes2 || distance2 != 10 {
		t.Errorf("EdgeExists failed")
	}

	no, distance3 := g.EdgeExists(v1, v4)
	if no || distance3 != 0 {
		t.Errorf("EdgeExists failed")
	}
}

// func TestShortestPath(t *testing.T) {
// 	g := Setup()
// 	g.Clear()

// 	if g.DistanceFrom(v0, v2) != 26 {
// 		t.Errorf("DistanceFrom failed")
// 	}

// 	// Returning 38 (1 to 0, 0 to 3, 3 to 4)
// 	// Shortest is 34 (1 to 0, 0 to 5, 5 to 4)
// 	if g.DistanceFrom(v1, v4) != 34 {
// 		t.Errorf("DistanceFrom failed")
// 	}
// }
