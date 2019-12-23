package main

import (
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

func main() {
	validate := validator.New()
	var a =0
	err := validate.Var(a,"required")
	if err != nil {
		fmt.Println(err)

	}
	var b bool
	err = validate.Var(b,"required")
	if err != nil {
		fmt.Println(err)
	}
	var s string =""
	err = validate.Var(s,"required")
	if err != nil {
		fmt.Println(err)
	}
}
