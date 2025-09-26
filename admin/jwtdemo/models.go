package main

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// --- API payloads ---

type LoginRequest struct {
	Username string `json:"username" example:"admin"`
	Password string `json:"password" example:"admin"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token" example:"<jwt>"`
}

type APIError struct {
	Message string `json:"message" example:"invalid credentials"`
}

type Profile struct {
	Username string `json:"username" example:"admin"`
	Role     string `json:"role" example:"demo-admin"`
	Issuer   string `json:"issuer" example:"jwtdemo"`
	Expires  string `json:"expires" example:"2025-09-26T13:00:00Z"`
}

// --- JWT claims ---

// RegisteredClaims mirrors jwt.RegisteredClaims but as an embedded type for examples.
type RegisteredClaims = jwt.RegisteredClaims

type TokenClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	RegisteredClaims
}

func TimePtr(t time.Time) *jwt.NumericDate {
	return jwt.NewNumericDate(t)
}
