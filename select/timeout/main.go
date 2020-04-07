package main

import (
	"fmt"
	"time"
)

func foo(ch chan<- string){
	time.Sleep(3*time.Second)
	ch <-"foo"
}
func bar(ch chan<- string){
	time.Sleep(2*time.Second)
	ch <-"bar"
}

func main() {
	f :=make(chan string)
	b :=make(chan string)
	go bar(b)
	go foo(f)
	// 会出现超时
	select {
	case r :=<-f:
		fmt.Println(r)
	case r :=<-b:
		fmt.Println(r)
	case <-time.After(time.Second):
		fmt.Println("time out")
	}
}
