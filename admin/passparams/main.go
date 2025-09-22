package main

import (
	"fmt"

	"github.com/cisco/admin/process"
)

func main() {

	response := process.SendMessage("127.0.0.1:80", "Send Data to Server")
	fmt.Println(response)

}
