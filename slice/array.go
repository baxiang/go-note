package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	newSlice := s[1:2:3]
	fmt.Printf("slice lenth=%d,cap=%d\n",len(s),cap(s))
	fmt.Printf("newSlice lenth=%d,cap=%d\n",len(newSlice),cap(newSlice))
}
