package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	//open the json file
	file, err := os.Open("users.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	//unmarshal the json file into a struct
	data := []map[string]interface{}{}
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		panic(err)
	}
	//print the struct
	for _, user := range data {
		for key, value := range user {
			fmt.Printf("%s: %v\n", key, value)
		}
	}

}
