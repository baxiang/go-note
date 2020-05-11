package service

import "errors"

type CalService interface {
	Add(a,b int)int //加
	Sub(a,b int)int //减
	Mul(a,b int)int //乘
	Div(a,b int)(int,error) //除
}

type Calculator struct {

}

func (c Calculator)Add(a,b int)int{
	return a+b
}

func (c Calculator)Sub(a,b int)int{
	return a-b
}

func (c Calculator)Mul(a,b int)int{
	return a*b
}

func (c Calculator)Div(a,b int)(int,error){
	if b == 0 {
		return 0, errors.New("被除数不能为零")
	}
	return a / b, nil
}