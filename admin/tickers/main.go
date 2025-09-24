package main

import (
	"time"
)

func HealthCheck(stop <-chan struct{}) {

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case t := <-ticker.C:
			println("Checking Database Health", t.String())
		case <-stop:
			println("Stopping Health Check")
			return
		}
	}
}

func main() {
	stop := make(chan struct{})
	go HealthCheck(stop)
	time.Sleep(10 * time.Second)
	close(stop)
	time.Sleep(5 * time.Second) // Wait to ensure goroutine has stopped

}
