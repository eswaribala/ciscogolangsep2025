package main

import (
	"fmt"

	"github.com/cisco/admin/process"
)

func init() {
	fmt.Println("Init function in main package...")
}



func main() {
	fmt.Println("Main function started...")

	process.CallTCP()
	CallTCP()
}
