package models

import (
	"encoding/csv"
	"errors"
	"os"
)

type Device struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Type        string `json:"type"`
	Network     struct {
		IPAddress  string `json:"ip_address"`
		MACAddress string `json:"mac_address"`
	}
}

//methods for Device struct can be added here

var deviceMap = make(map[string]*Device)

//Methods for Device struct

func (d *Device) Save() (bool, error) {
	deviceMap[d.ID] = d
	return true, nil
}

func (d *Device) Update() (bool, error) {
	if _, exists := deviceMap[d.ID]; exists {
		deviceMap[d.ID] = d
		return true, nil
	}
	return false, errors.New("device not found")
}

// Static methods or functions for Device struct

func FindAllDevices() ([]*Device, error) {
	devices := make([]*Device, 0, len(deviceMap))
	for _, device := range deviceMap {
		devices = append(devices, device)
	}
	return devices, nil
}

func FindDeviceByID(id string) (*Device, error) {
	if device, exists := deviceMap[id]; exists {
		return device, nil
	}
	return nil, errors.New("device not found")
}

func DeleteDeviceByID(id string) (bool, error) {
	if _, exists := deviceMap[id]; exists {
		delete(deviceMap, id)
		return true, nil
	}
	return false, errors.New("device not found")
}

func CreateCSVHeader(fileName string) (bool, error) {

	file, err := os.Create(fileName)
	if err != nil {
		return false, err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write([]string{"ID", "Name", "Description", "Type", "Status", "IP Address", "MAC Address"}); err != nil {
		return false, err
	}

	return true, nil
}

// methods to save device to CSV file
func (d *Device) SaveToCSV(fileName string) (bool, error) {

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return false, err
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	if err := writer.Write([]string{d.ID, d.Name, d.Description, d.Type, d.Status, d.Network.IPAddress, d.Network.MACAddress}); err != nil {
		return false, err
	}
	return true, nil
}
