package main

import (
	"errors"
	"fmt"
)

func operation(a,b int,symbol string)(result int,err error){
	switch symbol {
	case "+":
		return a+b,nil
	case "-":
		return a-b,nil
	case "*":
		return a*b,nil
	case "/":
		if b==0 {
			return 0,errors.New("division by zero")
		}
		return a/b,nil
	}
	return 0,errors.New("unsupported symbol:"+symbol)
}

func main() {
	r,_:=operation(2,3,"+")
	fmt.Println(r)
}
