package main

import (
	"fmt"
	"reflect"
)

type Bar struct {
	Foo string
	T *Test
}
type Test struct {
	A string
}

func main() {
	//b:=Bar{"hello"}
	//value := reflect.ValueOf(b)
	//fmt.Println(value.Type().Name())
	//value = reflect.Indirect(value)
	//
	//fmt.Println(value.Interface())
	//f:= value.FieldByName("Foo")
	//if f.Kind()==reflect.String&&f.CanSet(){
	//	f.SetString("world")
	//}
	//fmt.Println(f)
	//fmt.Println(value.Interface())

	b :=Bar{Foo:"world",T:& Test{
		A:"AA",
	}}
	v :=reflect.ValueOf(b)
	fmt.Println(v.Kind().String())
	fmt.Println(b)
	fmt.Println(v.CanAddr())
}
