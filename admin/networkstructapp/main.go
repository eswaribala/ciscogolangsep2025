package main

import (
	"fmt"

	"github.com/cisco/admin/models"
)

func main() {

	domain := models.NewDomain()
	fmt.Printf("Domain: %s\n", domain.ID)
	fmt.Printf("Name: %s\n", domain.Name)
	fmt.Printf("Description: %s\n", domain.Description)

	for _, subnet := range domain.Subnets {
		fmt.Printf("Subnet ID: %s, CIDR: %s, Gateway: %+v\n", subnet.ID, subnet.CIDR, subnet.GatewayInstance)
	}

}
