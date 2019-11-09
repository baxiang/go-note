package main

import (
	"fmt"
)
var a int
func foo()func (b int) int{
	return func(b int) int {
		fmt.Println(&a,a)
		a +=b
		return a
	}
}

func main() {
	f :=foo()
	fmt.Println(f(1))
	fmt.Println(f(1))
	g := foo()
	fmt.Println(g(1))
	fmt.Println(g(1))
}
