package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string `orm:"name"`
	Age int `orm:"age"`
}

func (s *Student)Update(){
	s.Age+=1
}
func (s Student)Show(){
	fmt.Println(s.Name,s.Age)
}
func (s Student)print(){
	fmt.Println(s.Name,s.Age)
}
func main() {
	s := Student{Name:"BX",Age:18}
	t :=reflect.TypeOf(s)
	for i:=0;i<t.NumField();i++{
		field := t.Field(i)
		fmt.Printf("name=%s PkgPath=%s index=%d type=%v tag=%v Offset=%v Anonymous=%v\n", field.Name,field.PkgPath, field.Index, field.Type, field.Tag.Get("orm"),field.Offset,field.Anonymous)
	}

	p :=reflect.TypeOf(&s)
	for i :=0;i<p.NumMethod();i++{
		 method := p.Method(i)
		fmt.Printf("name=%s PkgPath=%s index=%d type=%v  func=%v \n", method.Name,method.PkgPath, method.Index, method.Type, method.Func)
	}
}