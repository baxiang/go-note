package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now().UnixNano()
	fmt.Println(t)
	fmt.Println(t/ 1e6)
	fmt.Println(t/ 1e9)
	t = time.Now().Unix()
	fmt.Println(t)
}
