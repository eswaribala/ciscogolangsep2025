package models

import "fmt"

// DeviceNode represents a node in a hierarchical structure of devices.
//recursive structure
type DeviceNode struct {
	ID       int           `json:"id"`
	Name     string        `json:"name"`
	Location string        `json:"location"`
	Cost     float64       `json:"cost"`
	Children []*DeviceNode `json:"children"`
}

func TotalCost(node *DeviceNode) float64 {
	totalCost := node.Cost
	if node.ID == 1 {
		fmt.Printf("Node ID: %d, Cost: %.2f\n", node.ID, node.Cost)
	}
	for _, child := range node.Children {
		//recursion
		totalCost += TotalCost(child)
		fmt.Printf("Node ID: %d, Cost: %.2f\n", child.ID, child.Cost)
	}
	return totalCost
}
