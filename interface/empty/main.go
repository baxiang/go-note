package main

import "fmt"

type Foo interface {
	bar()
}
type FooType struct {

}

func (f FooType)bar(){
	fmt.Println("hello world")
}

func main() {
	var f Foo
	fmt.Println(f==nil)
	fmt.Printf("%T %v\n",f,f)

	var t *FooType
	fmt.Println(t==nil)
	fmt.Printf("%T %v\n",t,t)

	f = t
	fmt.Println(f==nil)
	fmt.Printf("%T %v\n",f,f)

}
