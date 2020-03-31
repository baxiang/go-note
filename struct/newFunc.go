package main

import "fmt"

type PersonNew struct {
	name string
}
type StudentNew struct {
	p     *PersonNew
	score float32
}

func main() {
	s := new(StudentNew)
	// fmt.Println(s.p.name) //会出现panic
	if s.p != nil {
		fmt.Println(s.p.name)
	}
}
