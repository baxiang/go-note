package main

import (
	"fmt"
)

func hello(c chan<- bool) {
	fmt.Println("hello goroutine")
	c <- true
}
func main() {
	c := make(chan bool)
	go hello(c)
	<-c
	fmt.Println("main goroutine")
}
