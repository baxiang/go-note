package main

import (
	"fmt"
	"reflect"
)

//判断实例是否实现了某接口
type Foo interface {
	show()
}

type Bar struct {
}
func (b *Bar) show() {
	fmt.Println("hello world")
}

func main() {
	b := new(Bar)
	f := reflect.TypeOf((*Foo)(nil)).Elem()
	tt := reflect.TypeOf(b)
	res := tt.Implements(f)
	fmt.Println(res)
}
