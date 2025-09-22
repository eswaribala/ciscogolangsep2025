package main

import (
	"fmt"

	"github.com/cisco/admin/process"
)

func init() {
	fmt.Println("Init function in main package...")
}

// order 2

func main() {
	fmt.Println("Main function started...")
	// order 3
	process.CallTCP()
	CallTCP()
}
