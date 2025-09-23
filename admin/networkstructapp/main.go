package main

import (
	"fmt"

	"github.com/cisco/admin/models"
)

func main() {
	
	domain := models.NewDomain()
	fmt.Printf("Domain: %+v\n", domain)

}
