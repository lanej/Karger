package main

import "fmt"

// Graph represents a collection of nodes
type Graph struct {
	Nodes map[int]*Node
}

// Node represents a node in the graph
type Node struct {
	Label int
	Graph *Graph
	Nodes map[int]*Node
}

// NewGraph initialize a new graph
func NewGraph() *Graph {
	var graph Graph
	graph.Nodes = make(map[int]*Node)
	return &graph
}

// Add a node to the graph
func (g *Graph) Add(label int) *Node {
	node := &Node{Label: label, Graph: g, Nodes: make(map[int]*Node)}
	g.Nodes[label] = node
	return node
}

// Relate two nodes
func (n *Node) Relate(label int) {
	relative := n.Graph.Nodes[label]

	if relative == nil {
		panic(fmt.Sprintf("Unknown node: %v", label))
	}

	relative.Nodes[n.Label] = n
	n.Nodes[relative.Label] = relative
}
