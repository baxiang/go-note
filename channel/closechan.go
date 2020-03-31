package main

import "fmt"

func sendMsg(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

func main() {
	c := make(chan int)
	go sendMsg(c)
	for {
		v, ok := <-c
		if ok == false {
			break
		}
		fmt.Print(v)
	}
}
