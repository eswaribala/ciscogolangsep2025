package models



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


