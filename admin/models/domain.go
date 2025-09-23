package models

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/mitchellh/mapstructure"
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

func DomainStructureToMap(domain *Domain) map[string]interface{} {
	//convert structure to map
	var domainMap map[string]interface{}
	mapstructure.Decode(domain, &domainMap)
	return domainMap

}
