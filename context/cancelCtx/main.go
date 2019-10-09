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
			return
		default:
			logger.Printf("work")
		}
	}
}
func main() {
	logger = log.New(os.Stdout, "cancel——Ctx", log.Ltime)
	ctx, cancel := context.WithCancel(context.Background())
	go downloadFile(ctx)
	time.Sleep(5 * time.Second)
	cancel()
	logger.Printf("cancel done")
}
