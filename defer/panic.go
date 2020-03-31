package main

import "fmt"

func Bar() {
	defer func() {
		fmt.Println("c")
	}()
	Foo()
	fmt.Println("继续执行d")
}

func Foo() {
	defer func() {
		if err :=recover();err!=nil{
			fmt.Println(err)
		}
		fmt.Println("b")
	}()
	panic("异常执行a")
}

func main(){
	Bar()
}