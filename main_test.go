package main

import "testing"

func TestGraphAdd(t *testing.T) {
	graph := NewGraph()
	node := graph.Add(3)

	if graph.Nodes[3] != node {
		t.Error("Failed to add node to graph")
	}
}

func TestNodeRelate(t *testing.T) {
	graph := NewGraph()
	node1 := graph.Add(3)
	node2 := graph.Add(6)

	node1.Relate(6)

	if node1.Nodes[6] != node2 {
		t.Error("Node 3 does not relate to Node 6")
	}

	if node2.Nodes[3] != node1 {
		t.Error("Node 6 does not relate to Node 3")
	}
}
