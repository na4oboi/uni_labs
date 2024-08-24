package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Node struct {
	Name     int
	Children []*Node
}

type Graph struct {
	nodes []*Node
}

func readGraph(filePath string) Graph {
	file, ferr := os.Open(filePath)
	if ferr != nil {
		panic(ferr)
	}

	scanner := bufio.NewScanner(file)
	graph := new(Graph)

	for scanner.Scan() {
		line := scanner.Text()
		nodes := strings.Split(line, " ")
		parentNode, errPr := strconv.Atoi(nodes[0])
		childNode, errCh := strconv.Atoi(nodes[1])
		if errPr != nil {
			panic(errPr)
		} else if errCh != nil {
			panic(errPr)
		}

		var parentNodeIndex int
		var childNodeIndex int

		if !GraphContains(graph.nodes, parentNode) {
			graph.nodes = append(graph.nodes, new(Node))
			parentNodeIndex = len(graph.nodes) - 1
			graph.nodes[parentNodeIndex].Name = parentNode
		} else {
			parentNodeIndex = GraphContainsIndex(graph.nodes, parentNode)
		}
		if !GraphContains(graph.nodes, childNode) {
			graph.nodes = append(graph.nodes, new(Node))
			childNodeIndex = len(graph.nodes) - 1
			graph.nodes[childNodeIndex].Name = childNode
		} else {
			childNodeIndex = GraphContainsIndex(graph.nodes, childNode)
		}

		graph.nodes[parentNodeIndex].Children = append(graph.nodes[parentNodeIndex].Children, graph.nodes[childNodeIndex])
	}

	return *graph
}

func (graph Graph) printGraphDFS(startNode int) {

	startNodeIndex := GraphContainsIndex(graph.nodes, startNode)
	startNodePointer := graph.nodes[startNodeIndex]

	stack := []*Node{startNodePointer}
	var visitedSlice []int

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		visitedSlice = append(visitedSlice, current.Name)
		stack = stack[:len(stack)-1]
		fmt.Printf("%d-", current.Name)

		for _, child := range current.Children {
			if !slices.Contains(visitedSlice, child.Name) {
				stack = append(stack, child)
			}
		}
	}
}

func GraphContains(sl []*Node, v int) bool {
	for _, vv := range sl {
		if vv.Name == v {
			return true
		}
	}
	return false
}

func GraphContainsIndex(sl []*Node, v int) int {
	for index, vv := range sl {
		if vv.Name == v {
			return index
		}
	}
	return -1
}
