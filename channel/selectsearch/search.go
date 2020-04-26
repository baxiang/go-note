package main

import (
	"fmt"
	"math/rand"
	"time"
)

func searchWord(engine string, c chan<- string, t time.Duration) {
	fmt.Printf("%s mock search time=%d\n", engine, t)
	time.Sleep(time.Second * t)
	c <- engine + " search result"
}

func main() {
	c := make(chan string)
	rand.Seed(time.Now().UnixNano())
	go searchWord("baidu", c, time.Duration(rand.Intn(10)))
	go searchWord("google", c, time.Duration(rand.Intn(10)))
	select {
	case r :=<-c:
		fmt.Println(r)
	case <-time.After(4*time.Second):
		fmt.Println("search time out")
	}
}
