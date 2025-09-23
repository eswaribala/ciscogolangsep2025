package models

type Device struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	IPAddress   string `json:"ip_address"`
	Description string `json:"description"`
	MACAddress  string `json:"mac_address"`
	Status      string `json:"status"`
	Type        string `json:"type"`
}