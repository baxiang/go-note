package main

import "fmt"

func main() {
	s:= []int{1:1}
	fmt.Printf("current slice point:%p\n" ,&s)
	fmt.Println("slice before function",s)
	modify(s)
	fmt.Println("slice after function",s)
}

func modify(s []int) {
	fmt.Printf("function slice point:%p\n" ,&s)
	s[1] = 2
}
