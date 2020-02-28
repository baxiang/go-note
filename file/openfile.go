package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//file, err := os.OpenFile("./TEST.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	file, err := os.Open("TEST.txt")
	if err!=nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()
	all, err := ioutil.ReadAll(file)
	if err!=nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(all))
}
