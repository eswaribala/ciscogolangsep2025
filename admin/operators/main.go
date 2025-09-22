package main

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v7"
)

func main() {

	//take random port number
	//loop
	var port int
	for i := 0; i < 10; i++ {
		port = gofakeit.IntRange(0, 65535)
		fmt.Println(port)

	}

}