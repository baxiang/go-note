package main

import "fmt"

type People struct{}

func (p *People) foo() {
	fmt.Println("people foo")
	p.bar()
}
func (p *People) bar() {
	fmt.Println("people bar")
}

type Student struct {
	 People
}

func (t *Student) bar() {
	fmt.Println("student bar")
}

func main() {
	t := Student{}
	t.foo()
}
