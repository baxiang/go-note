package main

import (
	"fmt"
	"time"
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
	time.Sleep(time.Second) // 为了演示随机输出 加一个延时
	// 会出现阻塞
	select {
	  case r :=<-f:
	  	fmt.Println(r)
	case r :=<-b:
		fmt.Println(r)
	}
}
