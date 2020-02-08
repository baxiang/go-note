package main

import (
	"fmt"
	"os"
)

func main() {
	file, e := os.Open("./test.txt")
	if e!=nil {
		fmt.Println(e.Error())
	}
	defer file.Close()
}
