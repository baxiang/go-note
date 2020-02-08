package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	//list := []int{7,2,3,4,5,6}
	//if len(list)>5 {
	//	list = list[1:4]// 左闭右开
	//}
	//fmt.Println(list)

	a := make([]int, 1, 3)
	//reflect.SliceHeader 为 slice运行时数据结构
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	fmt.Printf("%v\n",a)
	fmt.Printf("slice header: %#v\naddress of a: %p\n &a[0]: %p \n  &a: %p\n sh:%p \n",
		sh, a, &a[0],&a, sh)


}
