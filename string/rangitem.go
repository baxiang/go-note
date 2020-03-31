package main

import "fmt"

func main() {
	str := "Hello你好世界"
	fmt.Println(len(str))
	b := []byte(str)
	for _,v := range b{
		fmt.Print(string(v))
	}
	fmt.Println("")
	r :=[]rune(str)
	for _,v := range r{
		fmt.Print(string(v))
	}
	fmt.Println("")
	r[0]='h'
	r[5]='您'
	fmt.Println(string(r))
}
