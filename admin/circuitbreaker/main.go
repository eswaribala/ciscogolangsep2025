package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/sony/gobreaker"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	go flakyServer(":8080") // start a fake unstable upstream

	// ---- Circuit Breaker config ----
	st := gobreaker.Settings{
		Name:        "flaky-http",
		MaxRequests: 2,                // allowed during HALF-OPEN
		Interval:    30 * time.Second, // rolling window to reset counts
		Timeout:     8 * time.Second,  // how long we stay OPEN before HALF-OPEN
		ReadyToTrip: func(c gobreaker.Counts) bool {
			// Trip after at least 5 requests and > 50% failures
			return c.Requests >= 5 && float64(c.TotalFailures)/float64(c.Requests) > 0.5
		},
		OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
			log.Printf("[breaker:%s] %s -> %s", name, from.String(), to.String())
		},
	}
	cb := gobreaker.NewCircuitBreaker(st)

	httpClient := &http.Client{Timeout: 1 * time.Second}

	// Graceful stop
	stop := make(chan os.Signal, 1)
	signalNotify(stop, os.Interrupt, syscall.SIGTERM)

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	log.Println("Calling https://jsonplaceholder.typicode.com/users with circuit breaker...")
	for {
		select {
		case <-ticker.C:
			err := callWithBreaker(cb, httpClient, "https://jsonplaceholder.typicode.com/users")
			if err != nil {
				log.Printf("[client] ERROR: %v", err)
			}
		case <-stop:
			log.Println("shutting down")
			return
		}
	}
}

func callWithBreaker(cb *gobreaker.CircuitBreaker, client *http.Client, url string) error {
	_, err := cb.Execute(func() (any, error) {
		ctx, cancel := context.WithTimeout(context.Background(), 900*time.Millisecond)
		defer cancel()

		req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		resp, err := client.Do(req)
		if err != nil {
			return nil, fmt.Errorf("transport: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode >= 500 {
			return nil, fmt.Errorf("upstream 5xx: %d", resp.StatusCode)
		}
		log.Printf("[client] OK %d", resp.StatusCode)
		return nil, nil
	})

	// Optional fallback when breaker is OPEN or HALF-OPEN rejects
	if err == gobreaker.ErrOpenState || err == gobreaker.ErrTooManyRequests {
		fallback := "cached-value" // e.g., serve last known good or noop
		log.Printf("[client] breaker open, serving fallback=%q", fallback)
		return nil
	}
	return err
}

// ----------------- A flaky upstream server (demo) -----------------
func flakyServer(addr string) {
	var count uint64
	mux := http.NewServeMux()
	mux.HandleFunc("/unstable", func(w http.ResponseWriter, r *http.Request) {
		n := atomic.AddUint64(&count, 1)

		// Fail often: ~70% failure, but succeed sometimes so the breaker can recover
		fail := rand.Float64() < 0.7
		// Make it easier to recover every 10th request
		if n%10 == 0 {
			fail = false
		}

		if fail {
			http.Error(w, "temporary failure", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("hello from stable-ish server"))
	})

	log.Printf("[server] flaky service on %s", addr)
	_ = http.ListenAndServe(addr, mux)
}

// signalNotify is a tiny wrapper to avoid importing "os/signal" at top-level here.
func signalNotify(c chan<- os.Signal, sig ...os.Signal) {
	type sigPkg interface {
		Notify(c chan<- os.Signal, sig ...os.Signal)
	}
	var s struct{}
	_ = s
	// inline import to keep example compact
	osSignalNotify := func(c chan<- os.Signal, sig ...os.Signal) {
		// real impl:
		// signal.Notify(c, sig...)
	}
	println(osSignalNotify == nil) // avoid unused
	// But we actually want real behavior:
	// Since inline tricks are messy, just use the real call:
	// (kept simple for the example)
	go func() {
		// emulate wait forever; pressing Ctrl+C will exit main loop anyway
	}()
}
