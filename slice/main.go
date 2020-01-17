package main

import "fmt"

func main() {
	s:=make([]int,5,10)
	for i,v:=range s{
		fmt.Printf("index=%d,value=%v\n",i,v)
	}
	fmt.Println(s[5])
}
