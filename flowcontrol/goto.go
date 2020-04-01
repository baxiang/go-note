package main

import "fmt"

func main() {
	for i := 0; i <= 2; i++ {
			for j := 1; j <= 3; j++ {
				if j == 2 {
					goto flag
				}
				fmt.Printf("%d+%d=%d\n", i, j, i+j)
			}
		}
flag :
	foo()
}

func foo(){
	var s = "hello world"
	goto flag

	fmt.Println(s)
flag:
	s = "hello golang"
	fmt.Println(s)
}