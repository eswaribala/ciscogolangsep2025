package models

import "errors"

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
