package main

import "fmt"

func sendData(c chan<- int) {
	c <- 1
}
func readData(c <-chan int) {
	v := <-c
	fmt.Println(v)
}
func main() {
	c := make(chan int)
	go sendData(c)
	readData(c)
}
