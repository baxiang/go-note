package main

import "fmt"
func calc(idx string, a, b int) int {
	res := a + b
	fmt.Printf("[%s] %d+%d=%d\n",idx,a,b,res)
	return res
}

func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("2", a, b))
	a = 0
	defer calc("3", a, calc("4", a, b))
	b = 1
}

