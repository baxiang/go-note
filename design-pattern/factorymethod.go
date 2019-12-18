package main

import "fmt"

type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}


type OperatorBase struct {
	a,b int
}

func (o *OperatorBase)SetA(a int)  {
	o.a = a
}
func (o *OperatorBase)SetB(b int){
	o.b = b
}

type PlusOperator struct {
	*OperatorBase
}

func(o PlusOperator)Result()int {
	return o.a+o.b
}

type MinusOperator struct {
	*OperatorBase
}

func(o MinusOperator)Result()int {
	return o.a-o.b
}


type OperatorFactory interface {
	Create() Operator
}

type PlusOperatorFactory struct {

}

func(PlusOperatorFactory)Create()Operator{
	return &PlusOperator{&OperatorBase{}}
}


type MinusOperatorFactory struct {

}

func(MinusOperatorFactory)Create()Operator{
	return &MinusOperator{&OperatorBase{}}
}

func compute(factory OperatorFactory, a, b int) int {
	op := factory.Create()
	op.SetA(a)
	op.SetB(b)
	return op.Result()
}


func main() {
	fmt.Println(compute(PlusOperatorFactory{},2,1))
	fmt.Println(compute(MinusOperatorFactory{},2,1))
}
