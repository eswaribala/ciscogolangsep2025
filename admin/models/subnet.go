package models

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v7"
)

type Subnet struct {
	ID              string  `json:"id"`
	CIDR            string  `json:"cidr"`
	Description     string  `json:"description"`
	GatewayInstance Gateway  `json:"gateway_instance"`
}

func NewSubNetArray(count int) []*Subnet {
	subnets := make([]*Subnet, count)
	for i := 0; i < count; i++ {
		subnets[i] = &Subnet{
			ID:              fmt.Sprintf("subnet-%d", gofakeit.Number(1000, 9999)),
			CIDR:            fmt.Sprintf("192.168.%d.0/24", gofakeit.Number(0, 255)),
			Description:     fmt.Sprintf("Subnet %d", gofakeit.Name()),
			GatewayInstance: *NewGateway(),
		}
	}
	return subnets
}
