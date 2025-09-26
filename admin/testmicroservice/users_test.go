package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func fakeServer(payload any, status int) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/users" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_ = json.NewEncoder(w).Encode(payload)
	}))
}

func TestFetchUsers_OK(t *testing.T) {
	t.Parallel()
	in := []User{
		{ID: 1, Name: "Leanne Graham", Username: "Bret", Email: "Sincere@april.biz"},
		{ID: 2, Name: "Ervin Howell", Username: "Antonette", Email: "Shanna@melissa.tv"},
	}
	srv := fakeServer(in, http.StatusOK)
	defer srv.Close()

	svc := New(srv.URL, &http.Client{Timeout: time.Second})

	out, err := svc.FetchUsers(context.Background())
	require.NoError(t, err)
	require.Len(t, out, 2)
	require.Equal(t, "Bret", out[0].Username)
}

func TestFetchUsers_Non200(t *testing.T) {
	t.Parallel()
	srv := fakeServer(map[string]string{"error": "boom"}, http.StatusBadGateway)
	defer srv.Close()

	svc := New(srv.URL, &http.Client{Timeout: time.Second})
	_, err := svc.FetchUsers(context.Background())
	require.Error(t, err)
	require.Contains(t, err.Error(), "status 502")
}

func TestUniqueEmailDomains_AndHexID(t *testing.T) {
	t.Parallel()
	us := []User{
		{ID: 1, Email: "a@x.com"},
		{ID: 2, Email: "b@y.com"},
		{ID: 3, Email: "c@x.com"},
	}
	domains := UniqueEmailDomains(us)
	require.True(t, domains.Contains("x.com"))
	require.True(t, domains.Contains("y.com"))
	require.Equal(t, 2, domains.Cardinality())

	// go-ethereum hexutil usage
	require.Equal(t, "0x1", HexID(1))
	require.Equal(t, "0xa", HexID(10))
}
