package main

import (
	"fmt"
	"math/rand"
)

// Graph represents a collection of nodes
type Graph struct {
	Nodes map[int]*Node
	Edges []*Edge
}

// Edge represents to a connection between two nodes
type Edge struct {
	origin      *Node
	destination *Node
}

// Node represents a node in the graph
type Node struct {
	Label int
	Graph *Graph
	Edges []*Edge
}

// NewGraph initialize a new graph
func NewGraph() *Graph {
	var graph Graph
	graph.Nodes = make(map[int]*Node)
	return &graph
}

// Add a node to the graph
func (g *Graph) Add(label int) *Node {
	existing := g.Nodes[label]
	if existing != nil {
		return existing
	}
	node := &Node{Label: label, Graph: g}
	g.Nodes[label] = node
	return node
}

// Connect two nodes creating an Edge
func (g *Graph) Connect(a, b *Node) *Edge {
	if a == b {
		panic(fmt.Sprintf("Cannot add self-loop: %v", a.Label))
	}

	edge := &Edge{origin: a, destination: b}

	a.Edges = append(a.Edges, edge)
	b.Edges = append(b.Edges, edge)
	g.Edges = append(g.Edges, edge)

	return edge
}

// Reconnect edge
func (from *Node) Reconnect(to *Node, not *Edge) {
	for _, edge := range from.Edges {
		if edge == not {
			// do not process contracting edge
		} else if edge.destination == from {
			edge.destination = to
			to.Edges = append(to.Edges, edge)
		} else {
			edge.origin = to
			to.Edges = append(to.Edges, edge)
		}
	}
}

// Contract an edge
func (e *Edge) Contract() *Node {
	origin := e.origin
	destination := e.destination
	graph := origin.Graph

	supernode := graph.Add(origin.Label + destination.Label)

	origin.Reconnect(supernode, e)
	destination.Reconnect(supernode, e)

	delete(graph.Nodes, origin.Label)
	delete(graph.Nodes, destination.Label)

	var edgeIndex int
	for index, gEdge := range graph.Edges {
		if gEdge == e {
			edgeIndex = index
			break
		}
	}

	graph.Edges = append(graph.Edges[:edgeIndex], graph.Edges[(edgeIndex+1):]...)

	return supernode
}

// MinCut returns the size of the calculated MinCut using Krager's Random Contraction algorithm
func (g *Graph) MinCut() int {
	var index int
	var edge *Edge

	for len(g.Nodes) > 2 {
		index = rand.Intn(len(g.Edges))
		edge = g.Edges[index]

		edge.Contract()

		fmt.Printf("nodes remaining: %v, edges remaining: %v\n", len(g.Nodes), len(g.Edges))
	}

	return len(g.Edges)
}

func (g *Graph) print() {
	for index, edge := range g.Edges {
		fmt.Printf("Edge #%v: [%v, %v]\n", index, edge.destination.Label, edge.origin.Label)
	}
}

func (e *Edge) print() {
	fmt.Printf("Edge [%v, %v]\n", e.destination.Label, e.origin.Label)
}
