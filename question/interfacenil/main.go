package main

import (
	"fmt"
)

type People interface {
	Foo()
}

type Student struct{}

func (stu *Student) Foo() {

}

func newInstance() People {
	var stu *Student
	return stu
}

func main() {
	a := newInstance()
	if a == nil {
		fmt.Println("empty value")
	} else {
		fmt.Println(a)
		fmt.Printf("%p\n", a)
	}
}
