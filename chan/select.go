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
				fmt.Println("finish")
			default:
				fmt.Printf("wating")
			}
		}
	}()
	time.Sleep(10*time.Second)
	c<-true
	time.Sleep(5*time.Second)
}
