package main

import (
	"fmt"
	"shortest-way/model"
	"time"
)

const present = "present"

// findShortestPath accept list of all nodes in the graph, source and target nodes
// and returns the shortest path between source and target nodes, time which algorithm was working ang error
func findShortestPath(costs map[string]model.Costs, sourceNode string, targetNode string) (string, time.Duration, error) {

	//init
	var currentNode string
	var min int

	visited := map[string]bool{}
	d := map[string]int{}

	ver := []string{
		targetNode,
	}

	for nodeName := range costs {
		visited[nodeName] = false
		d[nodeName] = 1000000 //as some unreachable value
	}
	d[sourceNode] = 0

	//algorithm
	startTime := time.Now()

	for currentNode != present {
		currentNode = present
		min = 1000000

		for nodeName := range costs {
			if visited[nodeName] == false && d[nodeName] < min {
				min = d[nodeName]
				currentNode = nodeName
			}
		}

		if currentNode != "PRESENT" {
			for nodeName := range costs {
				if costs[currentNode][nodeName] > 0 {
					temp := min + costs[currentNode][nodeName]
					if temp < d[nodeName] {
						d[nodeName] = temp
					}
				}
			}
			visited[currentNode] = true
		}
	}

	weight := d[targetNode]

	for targetNode != sourceNode {
		for nodeName := range costs {
			if costs[nodeName][targetNode] != 0 { //if connection exist
				temp := weight - costs[nodeName][targetNode]
				if temp == d[nodeName] {
					weight = temp
					targetNode = nodeName
					ver = append(ver, nodeName)
				}
			}
		}
	}

	endTime := time.Now()

	//constructing output

	result := sourceNode
	for i := len(ver) - 1; i != 0; i-- {
		result += fmt.Sprintf(" -- %d --> %s", costs[ver[i]][ver[i-1]], ver[i-1])
	}

	return result, endTime.Sub(startTime), nil
}
