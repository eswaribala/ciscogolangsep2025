package main

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/cisco/admin/interfaces"
	"github.com/cisco/admin/models"
)

func main() {

	//call the interface

	var deviceDao interfaces.DeviceDAO
	status := []string{"active", "inactive", "maintenance"}

	dev := models.Device{
		ID:          gofakeit.UUID(),
		Name:        gofakeit.Name(),
		Description: gofakeit.Sentence(10),
		Type:        gofakeit.RandomString([]string{"router", "switch", "firewall", "access point"}),
		Status:      status[gofakeit.Number(0, 2)],
	}
	dev.Network.IPAddress = gofakeit.IPv4Address()
	dev.Network.MACAddress = gofakeit.MacAddress()

	//interface mapped to receiver
	deviceDao = &dev
	//call the method
	response, _ := deviceDao.Save()

	//print the response
	println(response)
}
