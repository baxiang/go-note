package main

import "fmt"

func main() {
	c := make(chan int, 3)
	fmt.Println(<-c)
}
