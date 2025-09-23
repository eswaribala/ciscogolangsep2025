package models

type Domain struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Subnets     []Subnet `json:"subnets"`
}
