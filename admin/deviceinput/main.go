package main

import (
	"fmt"
)

func main() {
	var deviceId string
	var deviceName string
	var IPAddress string
	var deviceType string

  _,err:= fmt.Scan(&deviceId, &deviceName, &IPAddress, &deviceType)
  if err != nil {
	fmt.Println("Error reading input:", err)
	return
  }

  fmt.Println("Device ID:", deviceId)
  fmt.Println("Device Name:", deviceName)
  fmt.Println("IP Address:", IPAddress)
  fmt.Println("Device Type:", deviceType)
}
