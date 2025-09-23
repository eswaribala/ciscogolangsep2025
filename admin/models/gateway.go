package models

import "github.com/brianvoe/gofakeit/v7"

type Gateway struct {
	ID          string `json:"id"`
	IPAddress   string `json:"ip_address"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Port        int    `json:"port"`
}

func NewGateway() *Gateway {
	gateway := &Gateway{
		ID:          gofakeit.UUID(),
		IPAddress:   gofakeit.IPv4Address(),
		Description: gofakeit.Sentence(5),
		Name:        gofakeit.Name(),
		Port:        gofakeit.Number(1024, 65535),
	}
	return gateway
}
