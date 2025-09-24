package main

import (
	"bufio"
	"os"
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

func monitorHealth(healthChannel chan string) {

	healthStatus := []string{"Healthy", "Degraded", "Unhealthy"}

	println("Starting Health Monitoring channel ...")

	// Simulate health monitoring
	for {
		time.Sleep(5 * time.Second)
		// In a real application, replace this with actual health check logic
		healthChannel <- gofakeit.RandomString(healthStatus) // Assume the service is healthy
	}
}
func monitorAlerts(alertChannel chan string) {
	alertMessages := []string{
		"High CPU Usage",
		"Memory Leak Detected",
		"Disk Space Low",
		"Network Latency High",
	}
	// Simulate alert monitoring
	for {
		time.Sleep(10 * time.Second)
		// In a real application, replace this with actual alert logic
		alertChannel <- gofakeit.RandomString(alertMessages) // Assume no alerts
	}
}
func monitorBandwidth(bandwidthChannel chan int) {
	bandwidthDataPercentage := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	// Simulate bandwidth monitoring
	for {
		time.Sleep(15 * time.Second)
		// In a real application, replace this with actual bandwidth check logic
		bandwidthChannel <- gofakeit.RandomInt(bandwidthDataPercentage) // Assume 100 Mbps
	}
}

func quit(quitChannel chan bool) {
	// Simulate quit signal after some time
	//time.Sleep(60 * time.Second)
	println("Press Enter to quit...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	quitChannel <- true
}

func main() {

	healthChannel := make(chan string)
	alertChannel := make(chan string)
	bandwidthChannel := make(chan int)
	quitChannel := make(chan bool)

	go monitorHealth(healthChannel)
	go monitorAlerts(alertChannel)
	go monitorBandwidth(bandwidthChannel)
	go quit(quitChannel)

	for {
		select {
		case healthStatus := <-healthChannel:
			println("Health Status Update:", healthStatus)
		case alertMessage := <-alertChannel:
			println("Alert Received:", alertMessage)
		case bandwidthUsage := <-bandwidthChannel:
			println("Bandwidth Usage Update:", bandwidthUsage, "%")
		case <-quitChannel:
			println("Quit signal received. Exiting...")
			return

		default:
			time.Sleep(1 * time.Second) // Prevent busy waiting
			println("No updates, system is stable...")
		}
	}

}
