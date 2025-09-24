package main

import (
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/cisco/admin/models"
)

func updateConfigWorker(updateConfigChannel chan *models.Device, devices *[]*models.Device) {
	count := 0
	time.Sleep(2 * time.Second)
	for _, dev := range *devices {
		updateConfigChannel <- dev
		count++
		println("Updating config for device:", count, "->", dev.Name)
		//time.Sleep(1 * time.Second) // Simulate time taken to update config

	}

}

func main() {
	//create the buffered channel
	updateConfigChannel := make(chan *models.Device, 5)
	status := []string{"active", "inactive", "maintenance"}
	count := gofakeit.Number(10, 15)

	devices := make([]*models.Device, count)

	for i := 0; i < count; i++ {
		dev := models.Device{
			ID:          gofakeit.UUID(),
			Name:        gofakeit.Name(),
			Description: gofakeit.Sentence(10),
			Type:        gofakeit.RandomString([]string{"router", "switch", "firewall", "access point"}),
			Status:      status[gofakeit.Number(0, 2)],
		}
		dev.Network.IPAddress = gofakeit.IPv4Address()
		dev.Network.MACAddress = gofakeit.MacAddress()
		devices[i] = &dev
	}

	//monitoring no of devices being updated

	go updateConfigWorker(updateConfigChannel, &devices)

	//simulate doing other work
	for i := 0; i < count; i++ {
		dev := <-updateConfigChannel
		println("Config updated for device:", dev.Name)
		time.Sleep(500 * time.Millisecond)
	}

	println("All device configurations updated.")
}
