package main

import "fmt"

// Graph represents a collection of nodes
type Graph struct {
	Nodes map[int]*Node
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

	return edge
}
