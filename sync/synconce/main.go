package main

import (
	"sync"
)

type Foo struct {
}

var(
	f *Foo
	once sync.Once
)
func NewFoo()*Foo{
	once.Do(func() {
		f = &Foo{}
	})
	return f
}

func main() {

}
