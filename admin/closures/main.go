package main

import (
	"fmt"
	"net/http"

	"github.com/brianvoe/gofakeit/v7"
)

func makeHandler(message string) http.HandlerFunc {
	//closure
	//anonymous function that captures the message variable
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", message+gofakeit.HackerPhrase())
	}
}

func main() {

	http.HandleFunc("/", makeHandler("API Ready!"))
	http.HandleFunc("/users", makeHandler("User Data"))
	http.HandleFunc("/products", makeHandler("Product Data"))
	http.HandleFunc("/orders", makeHandler("Order Data"))

	fmt.Println("Starting server on :7074")
	http.ListenAndServe(":7074", nil)
}
