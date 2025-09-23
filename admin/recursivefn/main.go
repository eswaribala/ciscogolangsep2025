package main

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/cisco/admin/models"
)

func main() {

	deviceNode := &models.DeviceNode{
		ID:       1,
		Name:     gofakeit.UUID(),
		Location: gofakeit.City(),
		Cost:     gofakeit.Price(1000, 5000),
	}

	for i := 2; i <= 20; i++ {
		childNode := &models.DeviceNode{
			ID:       i,
			Name:     gofakeit.UUID(),
			Location: gofakeit.City(),
			Cost:     gofakeit.Price(1000, 5000),
		}
		deviceNode.Children = append(deviceNode.Children, childNode)
	}

	totalCost := models.TotalCost(deviceNode)
	fmt.Printf("Total Cost of DeviceNode and its children: %f\n", totalCost)

}
