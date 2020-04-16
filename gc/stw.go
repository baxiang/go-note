package main

import (
	"fmt"
	"runtime"
	"time"
)
// 在<1.14 版本一直会无限卡死
func main() {
	go func() {
		for {
		}
	}()
	time.Sleep(time.Microsecond)
	runtime.GC()
	fmt.Println("finish")
}
