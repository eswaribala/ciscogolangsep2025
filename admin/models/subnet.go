package models

type Subnet struct {
	ID              string  `json:"id"`
	CIDR            string  `json:"cidr"`
	Description     string  `json:"description"`
	GatewayInstance Gateway `json:"gateway_instance"`
}
