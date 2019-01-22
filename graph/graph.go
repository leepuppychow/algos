package graph

import "fmt"

type Vertex struct {
	Data string
}

type Edge struct {
	V1       Vertex
	V2       Vertex
	Distance int
}

type Graph struct {
	AdjacencyList map[Vertex][]Vertex
	Visited       map[Vertex]bool
	Parents       map[Vertex]Vertex // key = vertex; value = parent
	EdgeList      []Edge
	VisitQueue    []Vertex
}

func Includes(collection []Vertex, v Vertex) bool {
	for _, vertex := range collection {
		if vertex == v {
			return true
		}
	}
	return false
}

func (g *Graph) Clear() {
	g.Visited = map[Vertex]bool{}
	g.VisitQueue = []Vertex{}
	g.Parents = map[Vertex]Vertex{}
}

func (g *Graph) Depth(v Vertex) int {
	depth := 0
	parent, ok := g.Parents[v]

	for ok {
		depth++
		parent, ok = g.Parents[parent]
	}
	return depth
}

func (g *Graph) EdgeExists(v1, v2 Vertex) (bool, int) {
	for _, edge := range g.EdgeList {
		// Undirected edges for now
		if (edge.V1 == v1 && edge.V2 == v2) || (edge.V1 == v2 && edge.V2 == v1) {
			return true, edge.Distance
		}
	}
	return false, 0
}

func (g *Graph) BreadthFirstSearch(start Vertex, searching Vertex) bool {
	g.Clear()
	g.VisitQueue = append(g.VisitQueue, start) // Push start node to VisitQueue

	for len(g.VisitQueue) > 0 { // Keep looping while there is stuff in the Queue
		current := g.VisitQueue[0] // Visit 1st node in VisitQueue and mark as visited (add to Visited)
		if current == searching {
			fmt.Println("****FOUND VERTEX****")
			return true
		}
		fmt.Println("CURRENT:", current)
		g.Visited[g.VisitQueue[0]] = true
		neighbors := g.AdjacencyList[current]
		for _, neighbor := range neighbors { // If current node has neighbors, check if they are in Visted or VisitQueue
			_, visited := g.Visited[neighbor]
			if visited || Includes(g.VisitQueue, neighbor) {
				continue
			} else {
				g.VisitQueue = append(g.VisitQueue, neighbor) // If not, push that neighbor to the VisitQueue
				g.Parents[neighbor] = current                 // neighbor's parent = current (store this to track path)
			}
		}
		g.VisitQueue = g.VisitQueue[1:] // Pop current node from Queue
		fmt.Println("VISIT QUEUE:", g.VisitQueue)
		fmt.Println("VISITED:", g.Visited)
		fmt.Println("PARENTS", g.Parents)
	}
	return false
}

func (g *Graph) DepthFirstSearch(current Vertex, searching Vertex) bool {
	fmt.Println("VISITED", g.Visited)
	g.Visited[current] = true
	if current == searching {
		fmt.Println("I FOUND YOU")
		return true
	}
	for _, neighbor := range g.AdjacencyList[current] {
		fmt.Println("CURRENT", current)
		fmt.Printf("NEIGHBOR %v; VISITED = %v\n", neighbor, g.Visited[neighbor])
		if !g.Visited[neighbor] {
			return g.DepthFirstSearch(neighbor, searching) // Was important to return here otherwise range loop would continue
		}
	}
	return false
}

func (g *Graph) ShortestPath(start, current Vertex) int {
	// TODO: Use Dijkstra's algorithm
	return 5
}
