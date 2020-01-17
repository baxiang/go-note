package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	wg sync.WaitGroup
	count int32
	mu sync.Mutex
)

func incCount() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		atomic.AddInt32(&count,1)
	}
}

func main() {
	wg.Add(2)
	go incCount()
	go incCount()
	wg.Wait()
	fmt.Println(count)
}
