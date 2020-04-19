package main

import (
	"fmt"
	"math/rand"
)
import "time"

func worker(id int, ready <-chan struct{}, done chan<- string) {
	<-ready // 阻塞在此，等待通知
	fmt.Println("Worker", id, "start")
	// 模拟一个工作负载。
	time.Sleep(time.Second * time.Duration(rand.Intn(10)))
	done <- fmt.Sprintf("Worker %d finish",id) // 数据回传
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ready, done := make(chan struct{}), make(chan string)
	for i :=1;i<=3;i++{
		go worker(i, ready, done)
	}
	// 模拟数据初始化过程
	time.Sleep(time.Second)
	// 单对多通知
	close(ready)
	index := 0
	for v :=range done{
		fmt.Println(v)
		index++
		if index==3 {
			close(done)
		}
	}
}
