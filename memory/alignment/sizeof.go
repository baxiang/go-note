package main

import (
	"fmt"
	"unsafe"
)

type bar struct {
	a int
	b string
}

type foo struct {
	a string
	b int
	c bool
	d string
}

func main() {
	fmt.Printf("string size: %d\n", unsafe.Sizeof("hello"))
	fmt.Printf("string size: %d\n", unsafe.Sizeof("hello go"))
	fmt.Printf("byte size: %d\n", unsafe.Sizeof(byte('h')))
	fmt.Printf("bool size: %d\n", unsafe.Sizeof(true))
	fmt.Printf("int8 size: %d\n", unsafe.Sizeof(int8(0)))
	fmt.Printf("int size: %d\n", unsafe.Sizeof(1))
	fmt.Printf("int32 size: %d\n", unsafe.Sizeof(int32(2)))
	fmt.Printf("int64 size: %d\n", unsafe.Sizeof(int64(3)))
	fmt.Printf("map size: %d\n", unsafe.Sizeof(map[string]int{"hello": 1}))
	fmt.Printf("list size: %d\n", unsafe.Sizeof([3]int{1, 2, 3}))
	fmt.Printf("list size: %d\n", unsafe.Sizeof([2]int{1, 2}))
	fmt.Printf("slice size: %d\n", unsafe.Sizeof([]int{1, 2, 3}))
	fmt.Printf("slice size: %d\n", unsafe.Sizeof([]int{1}))
	var s interface{}
	fmt.Printf("interface size: %d\n", unsafe.Sizeof(s))
	b := bar{a: 1, b: "2"}
	fmt.Printf("bar struct size: %d\n", unsafe.Sizeof(b))
	fmt.Printf("bar point size: %d\n", unsafe.Sizeof(&b))
	f := foo{a: "hello", b: 2, c: true, d: "world"}
	fmt.Printf("foo struct size: %d\n", unsafe.Sizeof(f))
	fmt.Printf("foo point size: %d\n", unsafe.Sizeof(&f))
	t := struct{}{}
	fmt.Printf("empty point size: %d\n", unsafe.Sizeof(t))
}
