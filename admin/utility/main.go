package main

import (
	"fmt"

	"github.com/bxcodec/faker/v4"
)

// global variable
var org string = "Cloud Academy"

func main() {
	fmt.Println("Kick Starting go language")
	fmt.Println("name", faker.FirstName())
	fmt.Println("email", faker.Email())
	fmt.Println("org", org)
	fmt.Println("org from func", getOrg())
	fmt.Println("org global", org)
	// accessing local variable from another function is not possible
	//fmt.Println("branch from func", branch);
}

func getOrg() string {
	// local variable
	var branch string = "DevOps"
	fmt.Println("branch", branch)
	// modifying the global variable
	org = "RPS Training"
	return org
}
