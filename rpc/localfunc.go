package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func main() {
	r := Add(2, 3)
	fmt.Println(r)
}
