package models

type Gateway struct {
	ID          string `json:"id"`
	IPAddress   string `json:"ip_address"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Port        int    `json:"port"`
}
