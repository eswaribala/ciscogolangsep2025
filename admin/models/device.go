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
