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
	// 1执行default 2 然后会随机打印<-f 或者<-b
	for {
		select {
		case r :=<-f:
			fmt.Println(r)
			return
		case r :=<-b:
			fmt.Println(r)
			return
		default:
			fmt.Println("default")
		}
	}
}