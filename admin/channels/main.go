package main

import (
	"log"
	"net/http"
)

func checkLink(link string, channel chan string) {

	println("Checking link:", link)
	resp, err := http.Get(link)

	if err != nil {
		log.Printf("Error checking %s: %v\n", link, err)

	}

	println(link + " - " + resp.Status)
	//send the message to the channel
	channel <- link + " - " + resp.Status + "-"

}

func main() {

	//create the channel
	channel := make(chan string)

	link := []string{"https://www.google.com", "https://www.facebook.com", "https://www.youtube.com", "https://www.twitter.com", "https://www.instagram.com"}

	var message string
	for _, url := range link {
		//rotines
		go checkLink(url, channel)
		//read the message from the channel
		message = <-channel
		println(message)

	}
	//listening to read all messages from the channel

	//println(<-channel)

	println("All links have been checked.")

}
