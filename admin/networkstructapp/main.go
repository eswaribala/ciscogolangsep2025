package main

import (
	"fmt"

	"github.com/cisco/admin/models"
)

func main() {

	domain := models.NewDomain()
	for key, value := range models.DomainStructureToMap(domain) {
		/*
			if key == "Subnets" {
				for i, subnet := range value.([]*models.Subnet) {
					for sk, sv := range models.SubnetStructureToMap(subnet) {
						fmt.Printf("Subnet %d - %s: %v\n", i+1, sk, sv)
					}

				}

			}
		*/
		fmt.Printf("%s: %v\n", key, value)

	}
}
