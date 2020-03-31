package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int,1)
	go func(c chan int){
		c<-1
		fmt.Println("你好, 世界")
		<-c
	}(c)
	time.Sleep(2*time.Second)
}

