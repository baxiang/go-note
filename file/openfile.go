package main

import (
	"log"
	"os"
)

func main() {
	file, e := os.Open("test.txt")
	if e!=nil {
		log.Fatal(e.Error())
	}
	defer file.Close()
}
