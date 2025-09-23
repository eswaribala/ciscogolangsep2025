package main

import (
	"github.com/brianvoe/gofakeit/v7"
)

func main() {

	devices := make([]string, 4)
	ports := make([]int, 4)

	for i := 0; i < 4; i++ {

		devices[i] = gofakeit.DomainName()
		ports[i] = gofakeit.IntRange(1, 2000)
	}

	for i, device := range devices {

		if ports[i] == 22 {
			println("SSH port detected!")
			continue
		}
		println("Device:", device, "Port:", ports[i])
	}
}
