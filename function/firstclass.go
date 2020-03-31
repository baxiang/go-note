package main

import "fmt"

func square(n int)int{
	return n*n
}
func negative(n int)int{
	return -n
}
func main() {
	f :=square
	fmt.Println(f(5))
	f =negative
	fmt.Println(f(5))
}
