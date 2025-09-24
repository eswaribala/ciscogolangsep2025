package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/cisco/admin/models"
)

// dedicated function to recover from panic
func recoverFromIPAddressPanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
}

func recoverFromAPIPanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
}

func isValidUUID(uuid string) bool {
	uuidRegex := `^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$`
	re := regexp.MustCompile(uuidRegex)
	return re.MatchString(uuid)
}

func isValidIPv4(ip string) bool {
	ipv4Regex := `^(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`
	re := regexp.MustCompile(ipv4Regex)
	return re.MatchString(ip)
}

func ValidateAPICall(url string) {
	defer recoverFromIPAddressPanic()

	resp, err := http.Get(url)
	if err != nil {
		panic("Invalid URL or network error: " + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic("Error: API call failed with status code ")

	}
	// Process the response
	// You can unmarshal the response body into a struct or perform other operations
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic("Error reading response body: " + err.Error())
	}
	println(string(body))

	println("API call successful")
}

func ValidateDevice(device *models.Device) bool {
	defer recoverFromAPIPanic()

	if device == nil {
		panic("device is nil")
	}
	if device.ID == "" || !isValidUUID(device.ID) {
		panic("device ID is empty or invalid")
	}

	if !isValidIPv4(device.Network.IPAddress) {
		panic("invalid IP address")
	}
	if device.Network.MACAddress == "" {
		panic("device MAC address is empty")
	}
	return true
}

func main() {
	status := []string{"active", "inactive", "maintenance"}

	dev := models.Device{
		ID:          gofakeit.UUID(),
		Name:        gofakeit.Name(),
		Description: gofakeit.Sentence(10),
		Type:        gofakeit.RandomString([]string{"router", "switch", "firewall", "access point"}),
		Status:      status[gofakeit.Number(0, 2)],
	}
	dev.Network.IPAddress = gofakeit.IPv4Address()
	dev.Network.MACAddress = gofakeit.MacAddress()
	// Add the Device to the map
	resp := ValidateDevice(&dev)

	if resp {
		fmt.Println("Device is valid")
	}

	ValidateAPICall("https://jsonplaceholder.typicode.com/users/11")

	fmt.Println("Main function completed.....status:", resp)
}
