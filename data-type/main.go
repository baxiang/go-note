package main

import "fmt"

func main() {
	str :="你好,世界"
	b := []byte(str)
	for _,v := range b{
		fmt.Print(string(v))
	}
	fmt.Print("\n")
	r :=[]rune(str)
	for _,v := range r{
		fmt.Print(string(v))
	}
	fmt.Print("\n")
	r[3]='未'
	r[4]='来'
	fmt.Println(string(r))
}
