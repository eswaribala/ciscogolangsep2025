package main

import (
	"sync"
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

func monitorHealth(wg *sync.WaitGroup) {

	healthStatus := []string{"Healthy", "Degraded", "Unhealthy"}

	println("Starting Health Monitoring channel ...")
	defer wg.Done()
	// Simulate health monitoring
	for i := 0; i < 3; i++ {
		time.Sleep(5 * time.Second)
		// In a real application, replace this with actual health check logic
		println("Health Status:", gofakeit.RandomString(healthStatus)) // Assume the service is healthy

	}
}
func monitorAlerts(wg *sync.WaitGroup) {
	alertMessages := []string{
		"High CPU Usage",
		"Memory Leak Detected",
		"Disk Space Low",
		"Network Latency High",
	}
	defer wg.Done()
	// Simulate alert monitoring
	for i := 0; i < 3; i++ {
		time.Sleep(10 * time.Second)
		// In a real application, replace this with actual alert logic
		println("Alert Received:", gofakeit.RandomString(alertMessages)) // Assume no alerts
	}
}
func monitorBandwidth(wg *sync.WaitGroup) {
	bandwidthDataPercentage := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	defer wg.Done()
	// Simulate bandwidth monitoring
	for i := 0; i < 3; i++ {
		time.Sleep(15 * time.Second)
		// In a real application, replace this with actual bandwidth check logic
		println("Bandwidth Usage:", gofakeit.RandomInt(bandwidthDataPercentage), "Mbps") // Assume 100 Mbps
	}

}

func main() {

	var wg sync.WaitGroup
	wg.Add(3)

	go monitorHealth(&wg)
	go monitorAlerts(&wg)
	go monitorBandwidth(&wg)

	wg.Wait()

	println("Work completed. Exiting...")

}
