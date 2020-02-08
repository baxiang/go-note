package main

import "fmt"

func main() {
	array := [3]int{1:2}
	modify(&array)
	fmt.Println("current list",array)
}

func modify(a *[3]int){
	a[1] =3
	fmt.Println("modify list",*a)
}
