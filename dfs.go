package main

import (
	"errors"
	"fmt"
)

func main() {

	// Random comment
	const graphFilePath = "graph.txt"
	const startNode = -1
	const finishNode = 3
	var inputErr = errors.New("can't use with negative numbers")

	if startNode > 0 && finishNode > 0 {
		graph := readGraph(graphFilePath)
		fmt.Println(graph.nodes)
		var pathLength int = graph.findPathLengthDFS(2, 3)
		fmt.Println()
		fmt.Printf("Path length from %d to %d is %d hops\n", startNode, finishNode, pathLength)
	} else {
		panic(inputErr)
	}

}
