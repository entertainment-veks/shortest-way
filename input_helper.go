package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"shortest-way/model"
	"strings"
)

// getAllNodeCosts accept the file path to json file with all nodes' info.
// and return slice of models, built on json-info in file
func getAllNodeCosts(nodesInfoFilePath string) (map[string]model.Costs, error) {
	data, err := os.ReadFile(nodesInfoFilePath)
	if err != nil {
		return nil, fmt.Errorf("unable to read data from nodes.json, err: %w", err)
	}

	var nodes []*model.Node
	if err := json.Unmarshal(data, &nodes); err != nil {
		return nil, fmt.Errorf("unable to unmarshal data from nodes.json, err: %w", err)
	}

	result := map[string]model.Costs{}

	for _, node := range nodes {
		result[node.Name] = node.Costs
	}

	return result, nil
}

// getTaskNodesNames waiting for user inputs source and target nodes, and return that
func getTaskNodesNames() (string, string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Source node is: ")
	sourceNode, err := reader.ReadString('\n')
	if err != nil {
		return "", "", fmt.Errorf("unable to read tast data from console, err: %w", err)
	}

	fmt.Print("Target node is: ")
	targetNode, err := reader.ReadString('\n')
	if err != nil {
		return "", "", fmt.Errorf("unable to read tast data from console, err: %w", err)
	}

	sourceNode = strings.Trim(sourceNode, "\n")
	targetNode = strings.Trim(targetNode, "\n")

	return sourceNode, targetNode, nil
}
