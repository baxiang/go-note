package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now().UnixNano()
	fmt.Println(t)
	t = time.Now().UnixNano()/ 1e6
	fmt.Println(t)
	t = time.Now().Unix()
	fmt.Println(t)

}
