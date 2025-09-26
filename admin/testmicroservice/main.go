package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	base := os.Getenv("JPH_BASE")
	if base == "" {
		base = "https://jsonplaceholder.typicode.com"
	}
	svc := New(base, &http.Client{Timeout: 5 * time.Second})

	r := gin.Default()
	r.GET("/users", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 4*time.Second)
		defer cancel()
		list, err := svc.FetchUsers(ctx)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, list)
	})
	_ = r.Run(":8080")
}
