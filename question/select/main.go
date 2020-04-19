package main

import (
	"fmt"
	"time"
)
func main() {
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"
	time.Sleep(time.Second)
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		fmt.Println(value)
	}
}

