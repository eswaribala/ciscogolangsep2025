package main

import (
	"fmt"
)

// custom type for devicetype
type DeviceType string

const (
	UnKnown DeviceType = "Unknown"
	Router  DeviceType = "Router"
	Switch  DeviceType = "Switch"
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
	fmt.Println("Enter device details (ID, Name, IP Address, Type) separated by spaces:")
	response, err := fmt.Scan(&deviceId, &deviceName, &IPAddress, &deviceType)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	if response != 4 {
		fmt.Println("Please provide exactly four values.")
		return
	}
	fmt.Println("Device ID:", deviceId)
	fmt.Println("Device Name:", deviceName)
	fmt.Println("IP Address:", IPAddress)
	dType := ParseDeviceType(deviceType)
	fmt.Println("Device Type:", dType)
}
