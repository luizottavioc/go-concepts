package main

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Cat struct {
	Id string `json:"id"`
	Url string `json:"url"`
	Width int `json:"width"`
	Height int `json:"height"`
}

var dirFiles string = "./files"
var file string = "json-file.json"
var path string = dirFiles + "/" + file

var apiUrl string = "https://api.thecatapi.com/v1/images/search?limit=10"

func getBodyApi() string {
	res, _ := http.Get(apiUrl)

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}

func getCatsByApi() []Cat {
	body := getBodyApi()

	var cats []Cat
	json.Unmarshal([]byte(body), &cats)

	return cats
}

func getCatsByFile() []Cat {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []Cat{}
		}

		log.Fatal(err)
	}

	var cats []Cat
	json.Unmarshal([]byte(f), &cats)

	return cats
}

func writeJsonFile(cats []Cat) bool {
	data, errM := json.Marshal(cats)
	if errM != nil {
		log.Fatal(errM)
	}

	errW := os.WriteFile(path, data, 0644)
	if errW != nil {
		log.Fatal(errW)
	}

	return true
}

func main() {
	catsFile := getCatsByFile()
	catsApi := getCatsByApi()
	
	cats := append(catsFile, catsApi...)

	writeJsonFile(cats)

	log.Println("Success!")
}