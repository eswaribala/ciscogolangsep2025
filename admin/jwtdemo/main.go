package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	docs "github.com/cisco/admin/jwtdemo/docs" // <-- replace with your module name
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           JWT Auth Demo API
// @version         1.0
// @description     Simple JWT authentication example with Swagger UI (Gin + swaggo).
// @contact.name    API Support
// @contact.email   you@example.com
// @BasePath        /
// @schemes         http

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter: Bearer {your JWT}

// health check (public)
// @Summary      Health check
// @Description  Simple health endpoint
// @Tags         public
// @Success      200 {string} string "ok"
// @Router       /health [get]
func health(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

// login issues a JWT (public)
// @Summary     Login (demo)
// @Description Use demo creds to get a JWT: username=admin, password=admin
// @Tags        auth
// @Accept      json
// @Produce     json
// @Param       creds body LoginRequest true "Credentials"
// @Success     200 {object} TokenResponse
// @Failure     401 {object} APIError
// @Router      /login [post]
func login(c *gin.Context) {
	var body LoginRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, APIError{Message: "invalid request"})
		return
	}

	// Demo validation only
	if body.Username != "admin" || body.Password != "admin" {
		c.JSON(http.StatusUnauthorized, APIError{Message: "invalid credentials"})
		return
	}

	// Create token valid for 15 minutes
	token, err := CreateToken(TokenClaims{
		Username: body.Username,
		Role:     "demo-admin",
		RegisteredClaims: RegisteredClaims{
			Issuer:    "jwtdemo",
			ExpiresAt: TimePtr(time.Now().Add(15 * time.Minute)),
		},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, APIError{Message: "token generation failed"})
		return
	}

	c.JSON(http.StatusOK, TokenResponse{AccessToken: token})
}

// protected endpoint
// @Summary     Get profile (protected)
// @Description Requires Bearer token
// @Tags        profile
// @Security    BearerAuth
// @Produce     json
// @Success     200 {object} Profile
// @Failure     401 {object} APIError
// @Router      /api/profile [get]
func profile(c *gin.Context) {
	claims, ok := GetClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, APIError{Message: "missing or invalid token"})
		return
	}

	resp := Profile{
		Username: claims.Username,
		Role:     claims.Role,
		Issuer:   claims.Issuer,
		Expires:  claims.ExpiresAt.Time.Format(time.RFC3339),
	}
	c.JSON(http.StatusOK, resp)
}

func main() {
	r := gin.Default()

	// Public
	r.GET("/health", health)
	r.POST("/login", login)

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	docs.SwaggerInfo.Schemes = []string{"http"} // or {"https"} if TLS
	docs.SwaggerInfo.BasePath = "/"

	// Protected group
	api := r.Group("/api", JWTMiddleware())
	{
		api.GET("/profile", profile)
	}

	log.Println("listening on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
