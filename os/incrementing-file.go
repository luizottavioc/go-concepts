package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var dirFiles string = "./files"
var file string = "creating-file.txt"
var path string = dirFiles + "/" + file

func main() {
	if _, err := os.Stat(dirFiles); os.IsNotExist(err) {
		if err := os.Mkdir(dirFiles, 0755); err != nil {
			log.Fatal(err)
		}
	}

	var f *os.File

	if _, err := os.Stat(path); os.IsNotExist(err) {
		f, err = os.Create(path)
		if err != nil {
			log.Fatal(err)
		}
	}else{
		f, err = os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}

	defer f.Close()

	dataFile, err := ioutil.ReadFile(path)
	if err != nil {
	   panic(err)
	}

	lines := strings.Split(string(dataFile), "\n")
	qtdLines := strconv.Itoa(len(lines))
	text := "Line " + qtdLines + "\n"

	n, err := f.WriteString(text)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Success! wrote %d bytes\n", n)
	f.Sync()
}