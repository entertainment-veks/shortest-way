package main

import (
	"shortest-way/model"
	"time"
)

const wayLegPattern = " -- %s --> %s" //costs, node_name

// findShortestPath accept list of all nodes in the graph, source and target nodes
// and returns the shortest path between source and target nodes, time which algorithm was working ang error
func findShortestPath(costs map[string]model.Costs, sourceNode string, targetNode string) (string, time.Duration, error) {
	startTime := time.Now()

	//init
	var currentNode string
	var min int

	v := map[string]int{}
	d := map[string]int{}

	ver := map[string]int{}
	ver[targetNode] = 0

	for nodeName := range costs {
		v[nodeName] = 1
		d[nodeName] = 1000000 //as some unreachable value
	}
	d[sourceNode] = 0

	//algorithm
	for currentNode != "PRESENT" {
		currentNode = "PRESENT"
		min = 1000000

		for nodeName := range costs {
			if v[nodeName] == 1 && d[nodeName] < min {
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
			v[currentNode] = 0
		}
	}

	weight := d[targetNode]

	for targetNode != sourceNode {
		for nodeName := range costs {
			if costs[nodeName][targetNode] != 0 { //if connection exist
				temp := weight - costs[nodeName][targetNode]
				if temp == d[nodeName] { // значит из этой вершины и был переход
					weight = temp         // сохраняем новый вес
					targetNode = nodeName // сохраняем предыдущую вершину
					ver[nodeName] = weight
				}
			}
		}
	}

	endTime := time.Now()
	return "", endTime.Sub(startTime), nil
}
