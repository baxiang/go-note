package main

import "fmt"

func foo() {
	i := 1
	defer fmt.Println(i)
	i++
	return
}

func main() {
	 foo()
}
