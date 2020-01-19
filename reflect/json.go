package main

import (
	"reflect"
	"fmt"
)

type User struct { // tags
	Id int64   `json:"id"`
	Name string `json:"name"`
	Gender bool    `json:"gender"`
}

func main() {
	u := User{10001, "ming", true}
	t := reflect.TypeOf(u)
	v := reflect.ValueOf(u)
	for i := 0; i < t.NumField(); i++ {
		fmt.Printf("%v=%v\n", t.Field(i).Tag.Get("json"),v.Field(i).Interface())
	}
}
