package main

import (
	"sync"

	"github.com/cisco/admin/models"
)

var mu sync.Mutex

// critical section
func CriticalSection(topicName string) *models.KafkaConfig {

	kafkaConfig := &models.KafkaConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   topicName,
		GroupID: "my-group",
		Enabled: true,
		Port:    9092,
	}
	return kafkaConfig
}

func BandWidthSystemProducer(wg *sync.WaitGroup) {

	defer wg.Done()
	defer println("Bandwidth System Producer Started....")
	mu.Lock()
	config := CriticalSection("bandwidth_topic")
	println("Bandwidth System Producer Config:", config.Topic)
	mu.Unlock()
	// Simulate work
	// time.Sleep(2 * time.Second)
	println("Bandwidth System Producer Finished.")

}

func AlertSystemProducer(wg *sync.WaitGroup) {
	defer wg.Done()
	defer println("Alert System Producer Started....")
	mu.Lock()
	config := CriticalSection("alert_topic")
	println("Alert System Producer Config:", config.Topic)
	mu.Unlock()
}

func DashboardConsumer(wg *sync.WaitGroup) {
	defer wg.Done()
	defer println("Dashboard Consumer Started....")
	mu.Lock()
	config := CriticalSection("dashboard_topic")
	println("Dashboard Consumer Config:", config.Topic)
	mu.Unlock()
}

func main() {

	var wg sync.WaitGroup
	wg.Add(3)

	go BandWidthSystemProducer(&wg)
	go AlertSystemProducer(&wg)
	go DashboardConsumer(&wg)

	wg.Wait()

	println("Producers and Consumers have completed their tasks.")

}
