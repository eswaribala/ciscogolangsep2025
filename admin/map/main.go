package main

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/cisco/admin/models"
)

func main() {

	// Create a new map
	device := make(map[int]models.Device)
	status := []string{"active", "inactive", "maintenance"}

	for i := 1; i <= 5; i++ {
		// Create a new Device instance
		dev := models.Device{
			ID:          gofakeit.UUID(),
			Name:        gofakeit.Name(),
			IPAddress:   gofakeit.IPv4Address(),
			Description: gofakeit.Sentence(10),
			MACAddress:  gofakeit.MacAddress(),
			Type:        gofakeit.RandomString([]string{"router", "switch", "firewall", "access point"}),
			Status:      status[gofakeit.Number(0, 2)],
		}
		// Add the Device to the map
		device[i] = dev
	}

	// Print the map
	for key, value := range device {
		println(key, value.Name, value.IPAddress, value.Type, value.Status)
	}

	CreateDomainMap()

}

func CreateDomainMap() {

	// Create a new map
	domain := make(map[int]models.Domain)

	for i := 1; i <= 5; i++ {
		// Create a new Domain instance
		dom := models.Domain{
			ID:          gofakeit.UUID(),
			Name:        gofakeit.DomainName(),
			Description: gofakeit.Sentence(10),
			Subnets:     models.NewSubNetArray(3),
		}
		// Add the Domain to the map
		domain[i] = dom
	}
	// Print the map
	for key, value := range domain {
		println(key, value.Name, value.Description)
		for _, subnet := range value.Subnets {
			println("  Subnet:", subnet.CIDR, subnet.Description)
		}
	}
}
