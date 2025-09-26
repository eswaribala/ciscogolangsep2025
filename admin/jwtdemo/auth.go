package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var hmacSecret = []byte("super-secret-key-change-me")

// CreateToken builds an HMAC-SHA256 JWT for given claims.
func CreateToken(c TokenClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(hmacSecret)
}

// JWTMiddleware validates "Authorization: Bearer <token>" and stores claims in context.
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authz := strings.TrimSpace(c.GetHeader("Authorization"))
		if authz == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, APIError{Message: "missing bearer token"})
			return
		}

		var raw string
		parts := strings.SplitN(authz, " ", 2)
		// Accept: "Bearer <token>" OR "<token>"
		if len(parts) == 2 && strings.EqualFold(parts[0], "Bearer") && strings.TrimSpace(parts[1]) != "" {
			raw = strings.TrimSpace(parts[1])
		} else if len(parts) == 1 {
			raw = parts[0] // naked token (Swagger entered without "Bearer ")
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, APIError{Message: "missing bearer token"})
			return
		}

		var claims TokenClaims
		token, err := jwt.ParseWithClaims(raw, &claims, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrTokenUnverifiable
			}
			return hmacSecret, nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, APIError{Message: "invalid token"})
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}

// GetClaims fetches TokenClaims from Gin context.
func GetClaims(c *gin.Context) (TokenClaims, bool) {
	v, ok := c.Get("claims")
	if !ok {
		return TokenClaims{}, false
	}
	claims, ok := v.(TokenClaims)
	return claims, ok
}
