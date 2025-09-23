package main

import (
	"fmt"

	"github.com/cisco/admin/models"
)

func main() {
	//create gateway instance
	gateway := models.NewGateway()
	//print gateway details
	fmt.Printf("Gateway ID: %s\n", gateway.ID)
	fmt.Printf("IP Address: %s\n", gateway.IPAddress)
	fmt.Printf("Description: %s\n", gateway.Description)
	fmt.Printf("Name: %s\n", gateway.Name)
	fmt.Printf("Port: %d\n", gateway.Port)

}
