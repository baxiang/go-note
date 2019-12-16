package main

import "fmt"

func main() {
	list := []int{7,2,3,4,5,6}
	if len(list)>5 {
		list = list[:4]
	}
	fmt.Println(list)
}
