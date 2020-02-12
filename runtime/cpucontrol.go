package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	var sum int
	var i int
	var wg sync.WaitGroup
	wg.Add(1)
	for{
		if i<100000{
			go func() {
				sum = sum+1
			}()
			i++
		}else {
			wg.Done()
			break
		}
	}
	wg.Wait()
    fmt.Println(sum)
}
