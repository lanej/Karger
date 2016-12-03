package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraphAdd(t *testing.T) {
	graph := NewGraph()
	node := graph.Add(3)

	if graph.Nodes[3] != node {
		t.Error("Failed to add node to graph")
	}
}

func TestGraphConnect(t *testing.T) {
	graph := NewGraph()
	node1 := graph.Add(3)
	node2 := graph.Add(6)

	edge := graph.Connect(node1, node2)

	if edge.origin != node1 {
		t.Errorf("Expected edge.origin = %v, actual = %v", node1, edge.origin)
	}

	if edge.destination != node2 {
		t.Errorf("Expected edge.destination = %v, actual = %v", node1, edge.destination)
	}

	assert.Contains(t, node1.Edges, edge)
	assert.Contains(t, node2.Edges, edge)
}

// func TestNodeContract(t *testing.T) {
// 	graph := NewGraph()
//
// 	node1 := graph.Add(1)
// 	node2 := graph.Add(2)
// 	node3 := graph.Add(3)
// 	node4 := graph.Add(4)
//
// }
