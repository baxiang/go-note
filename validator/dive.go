package main

import (
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type Container struct {
	Array []string          `validate:"required,gt=0,dive,required"`
	Map   map[string]string `validate:"required,gt=0,dive,keys,max=10,endkeys,required,max=100"`
}

func main() {
	validate := validator.New()
	var empty Container
	if err := validate.Struct(empty);err!=nil{
		fmt.Println(err.Error())
		fmt.Println()
	}
	vContainer := Container{Array:[]string{""},Map:map[string]string{"test > than 10": ""}}
	if err := validate.Struct(vContainer);err!=nil{
		fmt.Println(err.Error())
	}
}
