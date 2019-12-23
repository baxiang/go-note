package main

import (
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	"strings"
)

type MyStruct struct {
	String string `validate:"pre-we"`
}

func main() {
	validate := validator.New()
	validate.RegisterValidation("pre-we", ValidateMyVal)
	s := MyStruct{String: "wechat"}
	if err := validate.Struct(s);err != nil {
		fmt.Printf("Err(s):\n%+v\n", err)
	}
	s.String = "not wechat"
	if err := validate.Struct(s);err != nil {
		fmt.Println(err)
	}
}

func ValidateMyVal(fl validator.FieldLevel) bool {
	return strings.HasPrefix(fl.Field().String(),"we")
}