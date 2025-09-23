package main

import (
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/cisco/admin/models"
)

func main() {
	status := []string{"active", "inactive", "maintenance"}
	FindType(gofakeit.Name())
	FindType(gofakeit.Number(1, 100))
	FindType(gofakeit.Bool())
	FindType(gofakeit.RandomString([]string{"hello", "world"}))
	FindType(gofakeit.Float32Range(1.0, 100.0))
	FindType(gofakeit.Date())
	dev := models.Device{
		ID:          gofakeit.UUID(),
		Name:        gofakeit.Name(),
		Description: gofakeit.Sentence(10),
		Type:        gofakeit.RandomString([]string{"router", "switch", "firewall", "access point"}),
		Status:      status[gofakeit.Number(0, 2)],
	}
	dev.Network.IPAddress = gofakeit.IPv4Address()
	dev.Network.MACAddress = gofakeit.MacAddress()
	FindType(dev)

}

// empty interface example with type assertion
func FindType(i interface{}) {
	switch v := i.(type) {
	case int:
		println("int", v)
	case string:
		println("string", v)
	case bool:
		println("bool", v)
	case float32:
		println("float32", v)
	case time.Time:
		println("time", v.String())
	case models.Device:
		println("Device", v.Name, v.Type, v.Status)
	default:
		println("unknown")
	}
}
