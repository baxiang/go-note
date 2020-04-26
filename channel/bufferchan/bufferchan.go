package main

import (
	"fmt"
	"math/rand"
	"time"
)

func mockWebRequest(r chan<- int, duration time.Duration){
	time.Sleep(time.Second*duration)
	r<-rand.Intn(100)
}


func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	// 模拟3个网络请求
	c := make(chan int ,3)
	for i :=0;i<cap(c);i++{
		go mockWebRequest(c,time.Duration(i+1))
	}
	for i :=0;i<cap(c);i++{
		fmt.Println(<-c)
	}
    close(c)
	// 最终执行时间主要有耗时支持的请求决定
	fmt.Println("exec time", time.Since(start))
}
