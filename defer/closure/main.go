package main

import (
	"fmt"
)

func main() {
	a := bar()
	fmt.Printf("a1=%d\n",a(1))
	fmt.Printf("a2=%d\n",a(1))
	fmt.Printf("a3=%d\n",a(1))

	b := bar()
	fmt.Printf("b1=%d\n",b(1))
	fmt.Printf("b2=%d\n",b(1))
	fmt.Printf("b3=%d\n",b(1))
}


func bar()func(int)int{
	var i int
	return func(j int) int {
		fmt.Printf("point=%v i=%d j=%d\n",&i,i,j)
		i+=j
		return i
	}
}