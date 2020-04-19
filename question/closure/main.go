package main

import (
	"sync"
	"fmt"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	for i := 0; i < 2; i++ {
		fmt.Printf("1a=%p\n", &i)
		go func() {
			fmt.Printf("2a=%p\n", &i)
			fmt.Println("i: ", i)
			wg.Done()
		}()
	}
	wg.Wait()
}
