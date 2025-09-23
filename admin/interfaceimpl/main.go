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

	//create CSV file header

	csvResponse, err := models.CreateCSVHeader("devices.csv")
	if err != nil {
		println("Error creating CSV header:", err.Error())
	}
	println("CSV Header Creation Response:", csvResponse)

	//call the method
	response, _ := deviceDao.Save()

	//save to CSV
	csvResponse, err = deviceDao.SaveToCSV("devices.csv")
	if err != nil {
		println("Error saving to CSV:", err.Error())
	}
	println("Save to CSV Response:", csvResponse)

	//print the response
	println(response)

	//update the device
	dev.Description = gofakeit.Sentence(15)
	dev.Status = status[gofakeit.Number(0, 2)]
	dev.Type = gofakeit.RandomString([]string{"router", "switch", "firewall", "access point"})
	dev.Network.IPAddress = gofakeit.IPv4Address()
	dev.Network.MACAddress = gofakeit.MacAddress()
	response, _ = deviceDao.Update()
	println(response)

	//find all devices
	devices, _ := models.FindAllDevices()
	for _, d := range devices {
		println("Device ID:", d.ID, "Name:", d.Name, "Description:", d.Description, "Type:", d.Type, "Status:", d.Status, "IP:", d.Network.IPAddress, "MAC:", d.Network.MACAddress)
	}

	//find device by ID
	foundDevice, err := models.FindDeviceByID(dev.ID)
	if err != nil {
		println("Error:", err.Error())
	} else {
		println("Found Device ID:", foundDevice.ID, "Name:", foundDevice.Name, "Description:", foundDevice.Description, "Type:", foundDevice.Type, "Status:", foundDevice.Status, "IP:", foundDevice.Network.IPAddress, "MAC:", foundDevice.Network.MACAddress)
	}
	//delete device by ID
	deleted, err := models.DeleteDeviceByID(dev.ID)
	if err != nil {
		println("Error:", err.Error())
	} else {
		println("Deleted Device ID:", dev.ID, "Success:", deleted)
	}

}
