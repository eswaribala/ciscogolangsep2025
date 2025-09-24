package main

import (
	"regexp"

	"github.com/brianvoe/gofakeit/v7"
)

// custom type
type Validator func(string) bool

func assignIPAddress(ip string, ipValidator Validator) {
	if ipValidator(ip) {
		println("Valid IP address:", ip)
	} else {
		println("Invalid IP address:", ip)
	}
}

func generateIPAddress(ipaddressRange *[]string) {

	//IIF
	func(ipAddressRange *[]string) {
		println(gofakeit.RandomString(*ipAddressRange))
	}(ipaddressRange)
}

func main() {

	//lambdas code
	otp := func(min int, max int) int {
		return gofakeit.Number(min, max)
	}

	println(otp(1000, 9999))

	ipAddressRange := []string{"192.168.1.1", "192.168.1.254"}
	generateIPAddress(&ipAddressRange)

	ipAddressValidator := func(ip string) bool {
		ipv4Regex := `^(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`
		re := regexp.MustCompile(ipv4Regex)
		return re.MatchString(ip)
	}

	// passing function as parameter
	assignIPAddress(gofakeit.IPv4Address(), ipAddressValidator)
}
