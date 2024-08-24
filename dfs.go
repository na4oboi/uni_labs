package main

import (
	"fmt"
)

func main() {

	const graphFilePath = "graph.txt"
	graph := readGraph(graphFilePath)
	fmt.Println(graph.nodes)
	graph.printGraphDFS(2)
}
