package main

import (
	"fmt"
	"github.com/baxiang/go-note/errcode/origin/err"
)

func main() {
	invalidPara:=err.ERR_CODE_INVALID_PARAMS
	fmt.Println(int(invalidPara),invalidPara)
}
