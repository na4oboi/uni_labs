package main

import (
	"fmt"
)

func main() {

	const graphFilePath = "graph.txt"
	const startNode = 2
	const finishNode = 3
	graph := readGraph(graphFilePath)
	fmt.Println(graph.nodes)
	var pathLength int = graph.findPathLengthDFS(2, 3)
	fmt.Println()
	fmt.Printf("Path length from %d to %d is %d hops\n", startNode, finishNode, pathLength)
}
