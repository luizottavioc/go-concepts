package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var apiUser string = "https://random-data-api.com/api/v2/users?response_type=json"
var apiAddress string = "https://random-data-api.com/api/v2/addresses?response_type=json"
var apiAppliance string = "https://random-data-api.com/api/v2/appliances?response_type=json"
var urls []string = []string{apiUser, apiAddress, apiAppliance}

type jsonData interface {}
  
func main() {
	rand.Seed(time.Now().UnixNano())

	randomIndex := rand.Intn(len(urls))
	chosenUrl := urls[randomIndex]

	resp, err := http.Get(chosenUrl)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("- URL chosen: %s\n", chosenUrl)
	fmt.Printf("- Response status code: %d\n", resp.StatusCode)
	fmt.Printf("- Response body: %s\n", string(body))
}