package main

import (
	"fmt"

	"github.com/cisco/admin/process"
)

func main() {

	devices := []string{"Router", "Switch", "Firewall", "Access Point", "Server", "Client", "Modem", "Hub", "Bridge", "Repeater"}

	for _, device := range devices[1:3] {
		fmt.Println("Device:", device)
	}

	response := process.SendMessage("127.0.0.1:80", "Send Data to Server")
	fmt.Println(response)

	filteredDevices := process.FilterDevices(&devices)
	//fmt.Println("Filtered Devices:", filteredDevices)

	for _, device := range *filteredDevices {
		fmt.Println("Device:", device)
	}
}
