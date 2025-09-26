package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
)

func main() {
	godotenv.Load(".env") // load .env file if present

	broker := os.Getenv("broker")
	if broker == "" {
		broker = "localhost:9092"
	}

	topic := os.Getenv("topic")
	if topic == "" {
		topic = "site-events"
	}

	group := os.Getenv("group")
	if group == "" {
		group = "go-consumer-1"
	}

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{broker},
		Topic:          topic,
		GroupID:        group, // consumer group for auto rebalancing + offset commits
		MinBytes:       1,     // fetch floor
		MaxBytes:       10e6,  // fetch ceiling (~10MB)
		MaxWait:        500 * time.Millisecond,
		CommitInterval: time.Second, // how often to commit if using ReadMessage
	})
	defer r.Close()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	log.Printf("Consuming topic=%s broker=%s group=%s ...", topic, broker, group)

	for {
		// Use FetchMessage + CommitMessages for at-least-once AFTER processing
		m, err := r.FetchMessage(ctx)
		if err != nil {
			if ctx.Err() != nil {
				log.Println("Shutting down consumer...")
				return
			}
			log.Fatalf("fetch error: %v", err)
		}

		// --- Your processing here ---
		log.Printf("partition=%d offset=%d key=%s value=%s", m.Partition, m.Offset, string(m.Key), string(m.Value))
		if len(m.Headers) > 0 {
			for _, h := range m.Headers {
				log.Printf("  header %s=%s", h.Key, string(h.Value))
			}
		}
		// ----------------------------

		if err := r.CommitMessages(ctx, m); err != nil {
			log.Printf("commit error: %v", err)
		}
	}
}
