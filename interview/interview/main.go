package main

import (
	"fmt"
)

func producer(nums...int)<-chan int{
	c :=make(chan int)
	go func() {
		defer close(c)
		for _,n:=range nums{
			c<-n
		}
	}()
	return c
}

func work(inch <-chan int)<-chan int{
	c :=make(chan int)
	go func() {
		defer close(c)
		for n :=range inch{
			c<-n*2
		}
	}()
	return c
}


func main1() {
	p :=producer(1,2,3,4,5)
	res :=work(p)
	for v:=range res{
		fmt.Print(v)
	}
}
func main() {
	Validation()
}

