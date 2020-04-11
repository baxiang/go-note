package main

import "fmt"

func main() {
	done :=make(chan string)
	go func() {
		done  <-"hello chan"
	}()
	fmt.Println(<-done)
}
