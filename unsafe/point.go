package main

import (
	"fmt"
	"unsafe"
)
type user struct {
	name string
	age int
}

//func main() {
//	u:=user{
//		name: "张三",
//		age:  10,
//	}
//	fmt.Println(u)
//	pName:=(*string)(unsafe.Pointer(&u))
//	*pName="李四"
//	pAge:=(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&u))+unsafe.Offsetof(u.age)))
//	*pAge = 20
//	fmt.Println(u)
//}

func main() {
	i:= 10
	p := &i
	var pFloat *float64 = (*float64)(unsafe.Pointer(p))
	*pFloat = (*pFloat)+2
	fmt.Println(i)
}