package main

import (
	"net/http"
	"time"

	"log"

	_ "github.com/cisco/network/docs" // This line is necessary for go-swagger to find your docs
	"github.com/cisco/network/store"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Site API
// @version 1.0
// @description This is api service for managing Sites
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email parameswaribala@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:7072
// @BasePath /
func main() {
	store.VaultConnection()
	time.Sleep(30 * time.Second) // Wait for Vault to be ready
	db := store.MySQLConnectionHelper()
	store.GetTableInstance(db)
	res := store.EnsureTopic()
	if res != nil {
		log.Fatalf("failed to ensure kafka topic: %v", res)
	}
	println("Kafka topic ensured")
	store.InitKafkaWriters()
	//defer store.CloseKafkaWriters()

	mux := http.NewServeMux()

	mux.HandleFunc("POST /sites/v1.0", store.SaveSiteInfo)
	mux.HandleFunc("GET /sites/v1.0", store.GetAllSiteInfo)
	mux.HandleFunc("GET /sites/v1.0/{id}", store.GetSiteInfoByID)
	mux.HandleFunc("GET /sites/v1.0/kafka/{id}", store.PublishSiteInfoByID)
	mux.HandleFunc("PUT /sites/v1.0", store.UpdateSiteInfoByID)
	mux.HandleFunc("DELETE /sites/v1.0/{id}", store.DeleteSiteInfoByID)

	// Swagger UI served at /swagger/
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	// Your own handlers
	// mux.HandleFunc("/claims", claimsHandler)

	log.Println("Server running at http://localhost:7072")
	log.Fatal(http.ListenAndServe(":7072", mux))

}
