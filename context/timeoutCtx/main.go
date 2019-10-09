package main

import (
	"context"
	"log"
	"os"
	"time"
)

var logger *log.Logger

func downloadFile(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			logger.Printf("timeout done")
			return
		default:
			logger.Printf("work")
		}
	}
}
func main() {
	logger = log.New(os.Stdout, "timeoutCtx——", log.Ltime)
	//ctx, _ := context.WithTimeout(context.Background(),3*time.Second)//相对时间
	ctx, _ := context.WithDeadline(context.Background(),time.Now().Add(3*time.Second))//绝对时间
	go downloadFile(ctx)
	time.Sleep(5 * time.Second)
	logger.Printf("finsh")
}
