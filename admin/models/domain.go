package models

import (
	"github.com/brianvoe/gofakeit/v7"
)

type Domain struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Subnets     []*Subnet `json:"subnets"`
}

func NewDomain() *Domain {
	return &Domain{
		ID:          gofakeit.UUID(),
		Name:        gofakeit.Name(),
		Description: gofakeit.Sentence(5),
		Subnets:     NewSubNetArray(3),
	}
}
