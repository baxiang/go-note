package main

import (
	"fmt"
	"unsafe"
)
type user1 struct {
	b byte
	i int32
	j string
}

type user2 struct {
	b byte
	j string
	i int32
}


func main() {
	var u1 user1
	var u2 user2

	fmt.Println("u1 offset is ",unsafe.Offsetof(u1.b))
	fmt.Println("u1 offset is ",unsafe.Offsetof(u1.i))
	fmt.Println("u1 offset is ",unsafe.Offsetof(u1.j))

	fmt.Println("u2 offset is ",unsafe.Offsetof(u2.b))
	fmt.Println("u2 offset is ",unsafe.Offsetof(u2.j))
	fmt.Println("u2 offset is ",unsafe.Offsetof(u2.i))

	fmt.Println("u1 size is ",unsafe.Sizeof(u1))
	fmt.Println("u2 size is ",unsafe.Sizeof(u2))
}


