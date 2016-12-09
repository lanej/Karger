package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
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

func TestNodeContract1(t *testing.T) {
	graph := NewGraph()

	node1 := graph.Add(1)
	node2 := graph.Add(2)
	node3 := graph.Add(3)

	graph.Connect(node1, node2)
	graph.Connect(node2, node3)
	edge := graph.Connect(node3, node1)

	supernode := edge.Contract()
	assert.Equal(t, 4, supernode.Label, "Failed to add the labels together")

	var vertices []int

	for _, edge := range supernode.Edges {
		vertices = []int{edge.origin.Label, edge.destination.Label}
		sort.Ints(vertices)
		assert.Equal(t, []int{2, 4}, vertices, "Edge contraction failed")
	}
}

func TestNodeContact2(t *testing.T) {
	graph := NewGraph()

	node1 := graph.Add(1)
	node2 := graph.Add(2)
	node3 := graph.Add(3)
	node4 := graph.Add(4)

	graph.Connect(node1, node3)
	edge1 := graph.Connect(node2, node3)
	edge2 := graph.Connect(node4, node3)

	edge1.Contract()

	graph.print()

	edge2.Contract()

	graph.print()
}

func TestMinCut1(t *testing.T) {
	graph := NewGraph()

	node1 := graph.Add(1)
	node2 := graph.Add(2)
	node3 := graph.Add(3)

	graph.Connect(node1, node2)
	graph.Connect(node2, node3)
	graph.Connect(node3, node1)

	minCutSize := graph.MinCut()

	assert.Equal(t, 2, minCutSize, "Invalid MinCut result")
}

func TestMinCut2(t *testing.T) {
	graph := NewGraph()

	node1 := graph.Add(1)
	node2 := graph.Add(2)
	node3 := graph.Add(3)
	node4 := graph.Add(4)

	graph.Connect(node1, node3)
	graph.Connect(node2, node3)
	graph.Connect(node4, node3)

	minCutSize := graph.MinCut()

	assert.Equal(t, 1, minCutSize, "Invalid MinCut result")
}

func TestHomework(t *testing.T) {
	input := integersFromFile("mincut.txt")
	graph := NewGraph()

	for _, row := range input {
		node := graph.Add(row[0])
		for _, i := range row[1:] {
			relative := graph.Add(i)
			graph.Connect(node, relative)
		}
	}

	minCutSize := graph.MinCut()

	fmt.Printf("MinCut Size: %v", minCutSize)
}

func integersFromFile(s string) [][]int {
	inputBytes, err := ioutil.ReadFile(s)

	if err != nil {
		panic(err)
	}

	inputString := string(inputBytes[:])
	rawRows := strings.Split(inputString, "\n")
	data := make([][]int, len(rawRows)-1)

	for row, i := range rawRows {
		values := strings.Split(i, "\t")

		cols := []int{}

		for _, v := range values {
			j, err := strconv.Atoi(strings.TrimSpace(v))

			if err == nil {
				cols = append(cols, j)
			}
		}

		if len(cols) > 0 {
			data[row] = cols
		}
	}

	return data
}
