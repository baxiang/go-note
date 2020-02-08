package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var b = true
	fmt.Println("bool=",unsafe.Sizeof(b))
	var i int = 8
	fmt.Println("int=",unsafe.Sizeof(i))
	s := "hello go"
	fmt.Println("string=",unsafe.Sizeof(s))
	t := struct {}{}
	fmt.Println("empty struct=",unsafe.Sizeof(t))
	n := struct {
		Name string
		Age int
	}{Name:"wang",Age:10}
	fmt.Println("struct=",unsafe.Sizeof(n))
	fmt.Println("point=",unsafe.Sizeof(&n))

}
