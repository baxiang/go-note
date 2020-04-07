package main

import(
	"time"
	"fmt"
)

func foo(ch chan<- string){
	time.Sleep(time.Second)
	ch <-"foo"
}
func bar(ch chan<- string){
	time.Sleep(2*time.Second)
	ch <-"bar"
}

func main() {
	f :=make(chan string)
	b :=make(chan string)
	go foo(f)
	go bar(b)
	// 会出现阻塞
	select {
	  case r :=<-f:
	  	fmt.Println(r)
	case r :=<-b:
		fmt.Println(r)
	}
}
