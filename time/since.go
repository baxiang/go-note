package main

import (
	"fmt"
	"time"
)

func test( s time.Time) {
	sum := 0
	for i := 0; i < 100000000; i++ {
		sum++
	}
	elapsed := time.Since(s)
	fmt.Println("该函数执行完成耗时：", elapsed)
}
func main() {
	test(time.Now())
}
