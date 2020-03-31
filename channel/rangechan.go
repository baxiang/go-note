package main

import "fmt"

func SendData(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

func main() {
	c := make(chan int)
	go SendData(c)
	for v :=range c{
		fmt.Print(v)
	}
	fmt.Println("")
	fmt.Println("main finish")
}

