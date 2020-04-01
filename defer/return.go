package main

import "fmt"

var name string = "go"

//func myfunc() string {
//	defer func() {
//		name = "python"
//	}()
//	fmt.Printf("myfunc 函数里的name：%s\n", name)
//	return name
//}

func myfuncTwo() (s string) {
	s = name
	defer func() {
		s = "python"
	}()
	fmt.Printf("myfunc 函数里的name：%s\n", s)
	return s
}

func main() {
	myname := myfuncTwo()
	fmt.Printf("main 函数里的name: %s\n", name)
	fmt.Println("main 函数里的myname: ", myname)
}
