package main

import (
	"fmt"
	"math/rand"
)

// Graph represents a collection of nodes
type Graph struct {
	List map[int][]int
}

func (g *Graph) vertices() []int {
	var keys []int
	for key := range g.List {
		keys = append(keys, key)
	}
	return keys
}

func (g *Graph) randomVertices() (int, int) {
	// Make a list of verticies
	vertices := g.vertices()

	// Choose a random vertex
	v1Index := rand.Intn(len(vertices))
	v1 := vertices[v1Index]
	connectedV1 := g.List[v1]

	// Choose a random connected vertex
	v2Index := rand.Intn(len(connectedV1))
	v2 := g.List[v1][v2Index]

	return v1, v2
}

func (g *Graph) contract() {
	// Pick an edge
	v1, v2 := g.randomVertices()

	// Connect edge
	g.List[v1] = append(g.List[v1], g.List[v2]...)

	// Repoint neighbors
	for _, neighbor := range g.List[v2] {
		for i, v := range g.List[neighbor] {
			if v == v2 {
				g.List[neighbor][i] = v1
			}
		}
	}

	// Delete self-loops
	var selfless []int
	for _, v := range g.List[v1] {
		if v != v1 {
			selfless = append(selfless, v)
		}
	}
	g.List[v1] = selfless

	// Remove edge
	delete(g.List, v2)
}

func (g *Graph) karger() int {
	for len(g.vertices()) > 2 {
		g.contract()
	}

	firstVertex := g.vertices()[0]
	edges := g.List[firstVertex]

	return len(edges)
}

func (g *Graph) mincut(args ...int) int {
	min := len(g.List)
	var target *Graph

	var trials uint64
	if len(args) == 1 {
		trials = uint64(args[0])
	} else {
		trials = nChooseK(uint64(len(g.List)), uint64(2))
	}

	for i := uint64(0); i < trials; i++ {
		target = g.clone()
		cutSize := target.karger()
		if cutSize < min {
			min = cutSize
		}
	}

	return min
}

func (g *Graph) clone() *Graph {
	newList := make(map[int][]int)
	for key, val := range g.List {
		newList[key] = val
	}
	return &Graph{List: newList}
}

func (g *Graph) print() {
	fmt.Println("Graph :")
	for key := range g.List {
		fmt.Printf("%v - %v\n", key, g.List[key])
	}
}

func nChooseK(n, k uint64) uint64 {
	return factorial(n) / (factorial(k) * factorial(n-k))
}

func factorial(x uint64) uint64 {
	var result uint64
	if x < 1 {
		result = 1
	} else {
		result = x * factorial(x-1)
	}
	return result
}
