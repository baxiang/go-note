package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)
	go func() {
		for{
			select {
			case <-c:
				fmt.Println("game over")
				return
			default:
				fmt.Println("waiting")
				time.Sleep(1*time.Second)
			}
		}
	}()
	fmt.Println("start")
	time.Sleep(10*time.Second)
	fmt.Println("send message")
	c<-true
}
