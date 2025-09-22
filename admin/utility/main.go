package main

import (
	"fmt"

	"github.com/bxcodec/faker/v4"
)

func main() {
	fmt.Println("Kick Starting go language")
	fmt.Println("name", faker.FirstName())
	fmt.Println("email", faker.Email())
}
