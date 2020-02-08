package main

import (
	"fmt"
	"reflect"
)

type user struct {
	name string
	age int
}

func (u *user)update(name string,age int){
	u.name = name
	u.age = age
}
func(u *user)Show(){
	fmt.Printf("name=%s,age=%d\n",u.name,u.age)
}

func main() {
	u := user{
		name: "tony",
		age:  10,
	}
	v := reflect.ValueOf(&u)
	name := reflect.ValueOf(&u.name)
	name.Elem().SetString("wang")
	update := v.MethodByName("update")
	if update.IsValid() {
		args :=[]reflect.Value{reflect.ValueOf("ellen"),reflect.ValueOf(20)}
		update.Call(args)
	}


	show := v.MethodByName("Show")
	if show.IsValid() {
		show.Call([]reflect.Value{})
	}
}
