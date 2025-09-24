package main

import (
	"fmt"
	"log"
	"net/http"
)

func checkLink(link string) {

	println("Checking link:", link)
	resp, err := http.Get(link)

	if err != nil {
		log.Printf("Error checking %s: %v\n", link, err)

	}

	fmt.Println(link + " - " + resp.Status)

}

func main() {

	link := []string{"https://www.google.com", "https://www.facebook.com", "https://www.youtube.com", "https://www.twitter.com", "https://www.instagram.com"}

	for _, url := range link {
		//rotines
		go checkLink(url)

	}

	fmt.Println("All links have been checked.")

}
