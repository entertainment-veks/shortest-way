package main

import (
	"fmt"
	"log"
)

const nodesInfoFilePath = "nodes.json"

func main() {
	log.Println("Getting nodes info...")

	costs, err := getAllNodeCosts(nodesInfoFilePath)
	if err != nil {
		log.Panic(err)
	}

	log.Println("Nodes info was gotten")

	sourceNode, targetNode, err := getTaskNodesNames()
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Counting shortest way from [node %s] -to> [node %s]...\n", sourceNode, targetNode)

	pathInfo, time, err := findShortestPath(costs, sourceNode, targetNode)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Shortest way founded in %s\n", time)

	fmt.Printf("Shortest way is:\n%s", pathInfo)
}
