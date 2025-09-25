package main

import (
	"net/http"

	"log"

	_ "github.com/cisco/admin/siteapi/docs" // This line is necessary for go-swagger to find your docs
	"github.com/cisco/admin/siteapi/store"
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

	mux := http.NewServeMux()

	mux.HandleFunc("POST /sites/v1.0", store.SaveSiteInfo)
	mux.HandleFunc("GET /sites/v1.0", store.GetAllSiteInfo)
	mux.HandleFunc("GET /sites/v1.0/{id}", store.GetSiteInfoByID)
	mux.HandleFunc("PUT /sites/v1.0/{id}", store.UpdateSiteInfoByID)
	mux.HandleFunc("DELETE /sites/v1.0/{id}", store.DeleteSiteInfoByID)

	// Swagger UI served at /swagger/
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	// Your own handlers
	// mux.HandleFunc("/claims", claimsHandler)

	log.Println("Server running at http://localhost:7072")
	log.Fatal(http.ListenAndServe(":7072", mux))

}
