package main

import (
	"fmt"
	"unsafe"
)

type foo struct{
	a int
	b string

}
func main() {
	n := foo{a: 1, b: "hello"}
	niPointer := (*int)(unsafe.Pointer(&n))
	*niPointer = 2
	njPointer := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&n)) + unsafe.Offsetof(n.b)))
	*njPointer = "world"
	fmt.Println(n)
}
