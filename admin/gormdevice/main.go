package main

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/cisco/admin/gormdevice/interfaces"
	"github.com/cisco/admin/gormdevice/store"
)

func main() {

	//db := store.MySQLConnectionHelper()
	//store.GetTableInstance(db)

	//invoke interface
	var deviceDao interfaces.DeviceDAO

	device := store.Device{
		HostName:    gofakeit.Name(),
		Description: gofakeit.Sentence(10),
		IPAddress:   gofakeit.IPv4Address(),
		Location:    gofakeit.City(),
		Status:      gofakeit.Bool(),
		CreatedAt:   "2025-09-01",
		UpdatedAt:   "2025-09-21",
	}

	deviceDao = &device

	//create a device
	response, err := deviceDao.CreateDevice()
	if err != nil {
		println("Error creating device:", err.Error())
	}
	println("Create Device Response:", response)
	//get all devices
	devices, err := deviceDao.GetAllDevices()
	if err != nil {
		println("Error getting all devices:", err.Error())
	}
	println("Get All Devices Response:", devices)
	//get device by ID
	deviceByID, err := deviceDao.GetDeviceByID(response.DeviceID)
	if err != nil {
		println("Error getting device by ID:", err.Error())
	}
	println("Get Device By ID Response:", deviceByID)
	//update device
	updatedDevice, err := deviceDao.UpdateDevice(gofakeit.City(), gofakeit.Bool())
	if err != nil {
		println("Error updating device:", err.Error())
	}
	println("Update Device Response:", updatedDevice)
	//delete device
	/*
		deleteStatus, err := deviceDao.DeleteDevice(response.DeviceID)
		if err != nil {
			println("Error deleting device:", err.Error())
		}
		println("Delete Device Response:", deleteStatus)

	*/

}
