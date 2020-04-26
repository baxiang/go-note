package main

import (
	"fmt"
	"math/rand"
	"time"
)

func mockWebRequest(duration time.Duration)<-chan int{
	r := make(chan int)
	go func() {
		time.Sleep(time.Second*duration)
		r<-rand.Intn(100)
	}()
	return r
}

func AddResult(a,b int)int {
	fmt.Printf("a=%d b=%d\n",a,b)
	return a+b
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	// 模拟2个网络请求
	a := mockWebRequest(1)
	b := mockWebRequest(2)
	fmt.Println("result:",AddResult(<-a,<-b))
	// 最终执行时间主要有耗时支持的请求决定
	fmt.Println("exec time", time.Since(start))
}
