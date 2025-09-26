package main


import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/sony/gobreaker"
	"golang.org/x/time/rate"
)

type Sleeper func(time.Duration)

type Config struct {
	Client          *http.Client
	Limiter         *rate.Limiter
	Breaker         *gobreaker.CircuitBreaker
	MaxAttempts     int
	BaseBackoff     time.Duration
	Sleep           Sleeper // injected to avoid real sleeping in tests
	RetryOnStatusFn func(code int) bool
}

type Client struct{ cfg Config }

func New(cfg Config) *Client {
	if cfg.MaxAttempts < 1 {
		cfg.MaxAttempts = 1
	}
	if cfg.BaseBackoff <= 0 {
		cfg.BaseBackoff = 100 * time.Millisecond
	}
	if cfg.Sleep == nil {
		cfg.Sleep = time.Sleep
	}
	if cfg.RetryOnStatusFn == nil {
		cfg.RetryOnStatusFn = func(code int) bool { return code >= 500 }
	}
	return &Client{cfg: cfg}
}

// GetJSON performs: rate-limit → breaker → retries → GET url → decode into out
func (c *Client) GetJSON(ctx context.Context, url string, out any) error {
	// 1) rate limit
	if c.cfg.Limiter != nil {
		if err := c.cfg.Limiter.Wait(ctx); err != nil {
			return err
		}
	}

	// 2) circuit breaker wraps the whole attempt-sequence
	exec := func() (any, error) {
		var lastErr error
		backoff := c.cfg.BaseBackoff
		for attempt := 1; attempt <= c.cfg.MaxAttempts; attempt++ {
			req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
			resp, err := c.cfg.Client.Do(req)

			// network error: retryable
			if err != nil {
				lastErr = err
			} else {
				defer resp.Body.Close()
				if resp.StatusCode >= 400 {
					b, _ := io.ReadAll(resp.Body)
					// 4xx non-retry by default; 5xx retry
					if c.cfg.RetryOnStatusFn(resp.StatusCode) {
						lastErr = errors.New(resp.Status)
					} else {
						return nil, errors.New(string(b))
					}
				} else {
					dec := json.NewDecoder(resp.Body)
					if err := dec.Decode(out); err != nil {
						return nil, err
					}
					return out, nil
				}
			}

			// stop early if context cancelled
			if ctx.Err() != nil {
				return nil, ctx.Err()
			}
			// backoff and retry
			c.cfg.Sleep(backoff)
			backoff *= 2
		}
		return nil, lastErr
	}

	if c.cfg.Breaker != nil {
		_, err := c.cfg.Breaker.Execute(exec)
		return err
	}
	_, err := exec()
	return err
}
