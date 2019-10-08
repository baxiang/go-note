package main

import (
	"fmt"
	"time"
)
//https://www.yuque.com/baxiang/golang/goroutine/
func helloGo(){
	fmt.Println("hello golang")
}
//func main(){
//	helloGo()
//	fmt.Println("hello world")
//}
//func main(){
//	go helloGo()
//	fmt.Println("hello world")
//}

func main(){
	go helloGo()
	time.Sleep(time.Second)
	fmt.Println("hello world")
}
