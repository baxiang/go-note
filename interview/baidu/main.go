package main

import (
	"fmt"
	"sync"
)

func count(wg *sync.WaitGroup){
	defer wg.Done()
	mt.Lock()
	number=number+1
	fmt.Println(number)
	mt.Unlock()
}
var number = 0
var mt sync.Mutex
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go count(&wg)
	go count(&wg)
	wg.Wait()
}
