package main

import (
	"unsafe"
	"fmt"
)

type foo struct {
	myBool  bool
	myInt64 int64
	myInt   int32
}
 type bar struct{
	 myInt64 int64
	 myBool  bool
	 myInt   int32
 }


func main() {
	f := foo{}
	fmt.Println(unsafe.Sizeof(f.myBool))  // 1
	fmt.Println(unsafe.Sizeof(f.myInt64)) // 8
	fmt.Println(unsafe.Sizeof(f.myInt))   // 4
	fmt.Println(unsafe.Sizeof(f))         // 24

	fmt.Println(unsafe.Offsetof(f.myBool))  // 0
	fmt.Println(unsafe.Offsetof(f.myInt64)) // 8
	fmt.Println(unsafe.Offsetof(f.myInt))   // 16

    fmt.Println("========内存对齐================")
	b := bar{}
	fmt.Println(unsafe.Sizeof(b.myBool))  // 1
	fmt.Println(unsafe.Sizeof(b.myInt64)) // 8
	fmt.Println(unsafe.Sizeof(b.myInt))   // 4
	fmt.Println(unsafe.Sizeof(b))         // 16

	fmt.Println(unsafe.Offsetof(b.myBool))  // 8
	fmt.Println(unsafe.Offsetof(b.myInt64)) // 0
	fmt.Println(unsafe.Offsetof(b.myInt))   // 12
}