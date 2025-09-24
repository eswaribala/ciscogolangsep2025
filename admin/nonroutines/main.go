package main

import (
	"fmt"
	"net/http"
)

func checkLink(link string) (bool, error) {
	resp, err := http.Get(link)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	return resp.StatusCode == 200, nil
}

func main() {

	link := []string{"https://www.google.com", "https://www.facebook.com", "https://www.youtube.com", "https://www.twitter.com", "https://www.instagram.com"}

	for _, url := range link {
		isValid, err := checkLink(url)
		if err != nil {
			fmt.Printf("Error checking %s: %v\n", url, err)
			continue
		}
		if isValid {
			fmt.Printf("%s is valid.\n", url+" - 200 OK")
		} else {
			fmt.Printf("%s is not valid.\n", url)
		}
	}

	fmt.Println("All links have been checked.")

}
