package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Client interface {
	Do(*http.Request) (*http.Response, error)
}

type Service struct {
	BaseURL string
	Client  Client
	Timeout time.Duration
}

func New(base string, c Client) *Service {
	if httpC, ok := c.(*http.Client); ok && httpC.Timeout == 0 {
		httpC.Timeout = 5 * time.Second
	}
	return &Service{BaseURL: base, Client: c, Timeout: 5 * time.Second}
}

// FetchUsers fetches /users and decodes them.
func (s *Service) FetchUsers(ctx context.Context) ([]User, error) {
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, s.BaseURL+"/users", nil)
	res, err := s.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status %d", res.StatusCode)
	}
	var users []User
	if err := json.NewDecoder(res.Body).Decode(&users); err != nil {
		return nil, err
	}
	return users, nil
}

// UniqueEmailDomains shows golang-set in action.
func UniqueEmailDomains(us []User) mapset.Set[string] {
	s := mapset.NewSet[string]()
	for _, u := range us {
		for i := range u.Email {
			if u.Email[i] == '@' && i+1 < len(u.Email) {
				s.Add(u.Email[i+1:])
				break
			}
		}
	}
	return s
}

// HexID uses go-ethereum’s hexutil to format an ID.
func HexID(id int) string {
	return hexutil.EncodeUint64(uint64(id))
}
