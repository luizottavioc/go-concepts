package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Cat struct {
	Id string `json:"id"`
	Url string `json:"url"`
	Width int `json:"width"`
	Height int `json:"height"`
}

func (c Cat) String() {
	fmt.Printf("Cat (%s):\n", c.Id)
	fmt.Printf("- height: %d\n", c.Height)
	fmt.Printf("- width: %d\n", c.Width)
	fmt.Printf("- url access: %s\n", c.Url)
}

var apiUrl string = "https://api.thecatapi.com/v1/images/search?limit=10"

func main() {
	res, _ := http.Get(apiUrl)

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	if err != nil {
		log.Fatal(err)
	}

	var cats []Cat
	json.Unmarshal([]byte(body), &cats)

	for _, cat := range cats {
		cat.String()
		fmt.Println()
	}
}