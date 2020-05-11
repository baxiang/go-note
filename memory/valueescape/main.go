package main

import "fmt"

func bar()int{
	t :=2
	return t
}
func main() {
	x :=bar()
	fmt.Println(x)
	y :=x+1
	fmt.Println(y)
}
