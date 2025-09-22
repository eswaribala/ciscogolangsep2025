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
		if port > 0 && port <= 1027 {
			fmt.Println("Reserved Port")
		} else if port >= 1028 && port <= 10000 {
			fmt.Println("Application Port")
		} else {
			fmt.Println("Custom Port")
		}

	}

}
