package main

import (
	"fmt"
	"reflect"
)

type UserInfo struct {
	Name string
	Age  int
}

func main() {
	u := &UserInfo{}
	userType := reflect.TypeOf(u)
	elem := reflect.New(userType.Elem()).Elem()
	elem.FieldByName("Name").SetString("tony")
	elem.FieldByName("Age").SetInt(20)
	fmt.Println(elem.Field(0),elem.Field(1))
}
