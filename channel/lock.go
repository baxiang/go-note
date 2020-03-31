package main

import (
	"fmt"
	"sync"
)

var (
	c   = make(chan struct{}, 1) //容量为1的缓冲信道
	sum int
)

func increment(x int, wg *sync.WaitGroup) {
	c <- struct{}{}
	sum += x
	wg.Done()
	<-c
}

func main() {
	var wg sync.WaitGroup
	v := 100
	wg.Add(v)
	for i := 1; i <= v; i++ {
		go increment(i, &wg)
	}
	wg.Wait()
	fmt.Println(fmt.Sprintf("1-%d的和是：%d", v, sum))
}
