package main

import "fmt"
type Error struct {
	ErrCode int
	ErrMsg string
}

func NewError(code int, msg string) *Error {
	return &Error{ErrCode: code, ErrMsg: msg}
}


func (err *Error) Error() string {
	return err.ErrMsg
}

func ops(isSu bool)*Error{
	var error *Error
	if ! isSu{
		error =&Error{
			ErrCode: -1,
			ErrMsg:  "error",
		}
	}
	return error
}


func main() {
	err := ops(true)
	if err!= nil {
		fmt.Println(err)
	}else {
		fmt.Println("success")
	}
}
