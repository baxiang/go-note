package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := struct {}{}
	t := reflect.TypeOf(s)
	fmt.Println(t.Size(),t.Kind())
}
