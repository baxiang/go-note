package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx,cancel := context.WithCancel(context.Background())
	go func(c context.Context) {
		for{
			select {
			case <-c.Done():
				fmt.Println("game over")
				return
			default:
				fmt.Println("waiting")
				time.Sleep(1*time.Second)
			}
		}
	}(ctx)
	fmt.Println("start")
	time.Sleep(10*time.Second)
	fmt.Println("send message")
	cancel()
}
