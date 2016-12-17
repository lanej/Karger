package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSampleCase1(t *testing.T) {
	graph := homeworkGraph("testcase1.txt")

	mincut := graph.mincut()

	assert.Equal(t, 2, mincut)
}

func TestSampleCase2(t *testing.T) {
	graph := homeworkGraph("testcase2.txt")

	mincut := graph.mincut()

	assert.Equal(t, 2, mincut)
}

func TestAssignment(t *testing.T) {
	graph := homeworkGraph("mincut.txt")

	minCutSize := graph.mincut(19990)

	fmt.Printf("Mincut Size: %v\n", minCutSize)
}

func homeworkGraph(filename string) *Graph {
	input := integersFromFile(filename)

	return &Graph{List: input}
}

func integersFromFile(s string) map[int][]int {
	inputBytes, err := ioutil.ReadFile(s)

	if err != nil {
		panic(err)
	}

	inputString := string(inputBytes[:])
	rawRows := strings.Split(inputString, "\n")
	data := make(map[int][]int)

	for _, i := range rawRows {
		values := strings.Split(i, "\t")

		if len(values) < 2 {
			values = strings.Split(i, " ")
		}

		cols := []int{}

		for _, v := range values {
			j, err := strconv.Atoi(strings.TrimSpace(v))

			if err == nil {
				cols = append(cols, j)
			}
		}

		if len(cols) > 0 {
			data[cols[0]] = cols[1:]
		}
	}

	return data
}
