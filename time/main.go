package main

import (
	"fmt"
	"time"
)

func usePrecise(dur time.Duration) bool {
	return dur < time.Second || dur%time.Second != 0
}

func main() {
	fmt.Println(usePrecise(time.Second))
}
