package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"github.com/sony/gobreaker"
	"golang.org/x/time/rate"
)

func fastClient() *http.Client {
	return &http.Client{Timeout: 2 * time.Second}
}

func noSleep(time.Duration) {} // avoids slow tests

func newBreakerForTests() *gobreaker.CircuitBreaker {
	return gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:        "tbreaker",
		MaxRequests: 1, // 1 probe on half-open
		Timeout:     200 * time.Millisecond,
		Interval:    0, // no rolling window reset
		ReadyToTrip: func(c gobreaker.Counts) bool { return c.ConsecutiveFailures >= 3 },
	})
}

func TestRetry_500Then200_SucceedsOnSecondAttempt(t *testing.T) {
	var hits int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n := atomic.AddInt32(&hits, 1)
		if n == 1 {
			w.WriteHeader(500)
			return
		}
		w.WriteHeader(200)
		w.Write([]byte(`{"id":1,"title":"ok"}`))
	}))
	defer srv.Close()

	var out Post
	c := New(Config{
		Client:      fastClient(),
		Limiter:     nil,
		Breaker:     newBreakerForTests(),
		MaxAttempts: 3,
		BaseBackoff: 1, // practically no wait thanks to noSleep
		Sleep:       noSleep,
	})
	if err := c.GetJSON(context.Background(), srv.URL, &out); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Title != "ok" || hits != 2 {
		t.Fatalf("expected 2 hits and title ok; got hits=%d, title=%q", hits, out.Title)
	}
}

func TestNoRetryOn404(t *testing.T) {
	var hits int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&hits, 1)
		http.Error(w, "nope", 404)
	}))
	defer srv.Close()

	var out Post
	c := New(Config{
		Client:      fastClient(),
		Breaker:     newBreakerForTests(),
		MaxAttempts: 3,
		BaseBackoff: 1,
		Sleep:       noSleep,
	})
	err := c.GetJSON(context.Background(), srv.URL, &out)
	if err == nil {
		t.Fatal("expected error for 404")
	}
	if hits != 1 {
		t.Fatalf("expected no retries on 404; hits=%d", hits)
	}
}

func TestCircuitBreaker_OpensAfterFailures_ThenHalfOpenThenClose(t *testing.T) {
	var hits int32
	// First 3 calls fail (trip), then succeed.
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n := atomic.AddInt32(&hits, 1)
		if n <= 3 {
			w.WriteHeader(500)
			return
		}
		w.WriteHeader(200)
		w.Write([]byte(`{"id":2,"title":"recovered"}`))
	}))
	defer srv.Close()

	br := newBreakerForTests()
	c := New(Config{
		Client:      fastClient(),
		Breaker:     br,
		MaxAttempts: 1, // single attempt makes it easy to reason about failures
		BaseBackoff: 1,
		Sleep:       noSleep,
	})

	var out Post
	// 3 failures → breaker should open
	for i := 0; i < 3; i++ {
		_ = c.GetJSON(context.Background(), srv.URL, &out)
	}
	if st := br.State(); st != gobreaker.StateOpen {
		t.Fatalf("expected breaker open, got %v", st)
	}

	// While open, calls fail fast
	if err := c.GetJSON(context.Background(), srv.URL, &out); err == nil {
		t.Fatal("expected open breaker error")
	}

	// Wait for Timeout to elapse → half-open probe allowed
	time.Sleep(220 * time.Millisecond)
	if st := br.State(); st != gobreaker.StateHalfOpen {
		t.Fatalf("expected half-open, got %v", st)
	}

	// Half-open probe succeeds (server now returns 200) → close breaker
	if err := c.GetJSON(context.Background(), srv.URL, &out); err != nil {
		t.Fatalf("expected success on probe: %v", err)
	}
	if st := br.State(); st != gobreaker.StateClosed {
		t.Fatalf("expected closed after success, got %v", st)
	}
	if out.Title != "recovered" {
		t.Fatalf("unexpected body: %+v", out)
	}
}

func TestRateLimit_ThrottlesRequests(t *testing.T) {
	// Server always succeeds
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(`{"id":1,"title":"ok"}`))
	}))
	defer srv.Close()

	lim := rate.NewLimiter(1, 1) // 1 token/sec, burst=1
	c := New(Config{
		Client:      fastClient(),
		Limiter:     lim,
		Breaker:     nil,
		MaxAttempts: 1,
		Sleep:       noSleep,
	})

	start := time.Now()
	var out Post
	// Two sequential calls with limiter 1 r/s should take ~>=1s (second waits)
	if err := c.GetJSON(context.Background(), srv.URL, &out); err != nil {
		t.Fatalf("err1: %v", err)
	}
	if err := c.GetJSON(context.Background(), srv.URL, &out); err != nil {
		t.Fatalf("err2: %v", err)
	}
	elapsed := time.Since(start)
	if elapsed < 900*time.Millisecond {
		t.Fatalf("expected throttling; elapsed=%v < 900ms", elapsed)
	}
}
