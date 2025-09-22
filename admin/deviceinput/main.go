package main

import (
	"bufio"
	"fmt"
	"os"
)

// custom type for devicetype
type DeviceType string

const (
	UnKnown  DeviceType = "Unknown"
	Router   DeviceType = "Router"
	Switch   DeviceType = "Switch"
	Firewall DeviceType = "Firewall"
	Server   DeviceType = "Server"
	Gateway  DeviceType = "Gateway"
)

func ParseDeviceType(input string) DeviceType {
	switch input {
	case "Router":
		return Router
	case "Switch":
		return Switch
	case "Firewall":
		return Firewall
	case "Server":
		return Server
	case "Gateway":
		return Gateway
	default:
		return UnKnown
	}
}

func main() {
	var deviceId string
	var deviceName string
	var IPAddress string
	var deviceType string
	fmt.Println("Enter device details (ID, IP Address, Type) separated by spaces:")
	response, err := fmt.Scan(&deviceId, &IPAddress, &deviceType)
	fmt.Println("Response:", response)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	// Flush leftover newline
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	fmt.Print("Enter Device Name: ")
	deviceName, _ = reader.ReadString('\n')
	fmt.Println("Device ID:", deviceId)
	fmt.Println("Device Name:", deviceName)
	fmt.Println("IP Address:", IPAddress)
	dType := ParseDeviceType(deviceType)
	fmt.Println("Device Type:", dType)
}
