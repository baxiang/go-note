package main

import (
	"fmt"
)

func foo(ch chan<- string){
	ch <-"foo"
}
func bar(ch chan<- string){

	ch <-"bar"
}

func main() {
	f :=make(chan string)
	b :=make(chan string)
	go bar(b)
	go foo(f)
	// 执行默认结果
	select {
	case r :=<-f:
		fmt.Println(r)
	case r :=<-b:
		fmt.Println(r)
	 default:
		fmt.Println("default")
	}
}