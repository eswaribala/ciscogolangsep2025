package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/sony/gobreaker"
	"golang.org/x/time/rate"
)

// ---- Domain types ----
type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// ---- Config ----
const (
	baseURL           = "https://jsonplaceholder.typicode.com"
	apiPath           = "/posts/1"
	clientTimeout     = 4 * time.Second
	retryMaxAttempts  = 3
	retryBaseBackoff  = 200 * time.Millisecond
	ratePerSecond     = 2 // average RPS
	rateBurst         = 5 // allow short bursts
	breakerFailThresh = 5 // trip after N consecutive failures
	breakerOpenFor    = 5 * time.Second
	breakerWindow     = 30 * time.Second
)

// ---- HTTP client with sane timeouts ----
var httpClient = &http.Client{
	Timeout: clientTimeout,
	Transport: &http.Transport{
		TLSHandshakeTimeout:   3 * time.Second,
		ResponseHeaderTimeout: 3 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		IdleConnTimeout:       30 * time.Second,
		MaxIdleConns:          100,
		DialContext: (&net.Dialer{
			Timeout:   3 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
	},
}

// ---- Rate limiter (token bucket) ----
var limiter = rate.NewLimiter(rate.Limit(ratePerSecond), rateBurst)

// ---- Circuit breaker ----
var breaker *gobreaker.CircuitBreaker

func init() {
	settings := gobreaker.Settings{
		Name:        "jsonplaceholderBreaker",
		MaxRequests: 1,              // during half-open, allow 1 probe
		Interval:    breakerWindow,  // reset counts every window
		Timeout:     breakerOpenFor, // stay open this long
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			// Trip after consecutive failures threshold OR high error ratio
			if counts.ConsecutiveFailures >= breakerFailThresh {
				return true
			}
			// Optional: trip if error rate > 50% after at least 10 calls
			total := counts.Requests
			fail := counts.TotalFailures
			return total >= 10 && float64(fail)/float64(total) > 0.5
		},
	}
	breaker = gobreaker.NewCircuitBreaker(settings)
}

// ---- Resilient call: rate-limit -> breaker -> retry/backoff -> request ----
func getPost(ctx context.Context) (Post, error) {
	var result Post

	// 1) Rate limit (wait for a token respecting ctx deadline)
	if err := limiter.Wait(ctx); err != nil {
		return result, fmt.Errorf("rate limit wait: %w", err)
	}

	// 2) Circuit breaker wraps the whole attempt (including retries).
	//    We consider "success" only if one retry succeeds; otherwise breaker sees a failure.
	out, err := breaker.Execute(func() (any, error) {
		var lastErr error
		backoff := retryBaseBackoff

		for attempt := 1; attempt <= retryMaxAttempts; attempt++ {
			// respect ctx for each attempt
			req, _ := http.NewRequestWithContext(ctx, http.MethodGet, baseURL+apiPath, nil)

			resp, err := httpClient.Do(req)
			if err == nil && resp != nil && resp.StatusCode >= 500 {
				// Treat 5xx as transient (retryable)
				err = fmt.Errorf("server %d", resp.StatusCode)
			}

			// Handle response
			if err == nil && resp != nil && resp.StatusCode < 500 {
				defer resp.Body.Close()
				if resp.StatusCode >= 400 {
					// 4xx: treat as non-retryable (client issue)
					b, _ := io.ReadAll(resp.Body)
					return nil, fmt.Errorf("non-retryable status %d: %s", resp.StatusCode, string(b))
				}

				dec := json.NewDecoder(resp.Body)
				var p Post
				if err := dec.Decode(&p); err != nil {
					return nil, fmt.Errorf("decode: %w", err)
				}
				return p, nil // success
			}

			// transient failure path: prepare retry
			if resp != nil {
				resp.Body.Close()
			}
			lastErr = err

			// if context already done, abort early (don’t spin)
			if ctx.Err() != nil {
				break
			}

			// exponential backoff with jitter
			sleep := backoff + time.Duration(randJitterMillis(60))*time.Millisecond
			timer := time.NewTimer(sleep)
			select {
			case <-ctx.Done():
				timer.Stop()
				return nil, ctx.Err()
			case <-timer.C:
			}
			backoff *= 2
		}
		return nil, fmt.Errorf("all retries failed: %w", lastErr)
	})
	if err != nil {
		return result, err
	}

	// type assert the successful value from breaker
	p, ok := out.(Post)
	if !ok {
		return result, errors.New("unexpected type from breaker")
	}
	return p, nil
}

// tiny jitter helper (no extra deps)
func randJitterMillis(max int64) int64 {
	// not crypto rand; fine for backoff jitter
	return time.Now().UnixNano()%max + 1
}

func main() {
	fmt.Println("Calling JSONPlaceholder with resiliency...")

	// Example: fire a few calls in a loop to see limiter + breaker behavior.
	for i := 1; i <= 8; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		p, err := getPost(ctx)
		if err != nil {
			// When breaker is open you’ll see: gobreaker: circuit breaker is open
			fmt.Printf("[%d] ERROR: %v\n", i, err)
			continue
		}
		fmt.Printf("[%d] OK: Post #%d - %q\n", i, p.ID, p.Title)
	}
}
