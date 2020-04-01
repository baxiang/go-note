package main

import (
	"fmt"
	"time"
)

func foo(x int,c chan int){

	c <- x


}
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	go Getdata(1,6,ch1)
	go Getdata(2,3,ch2)
	go Getdata(3,5,ch3)
	select{
	case v:=<- ch1:
		fmt.Println(v)
	case v:=<- ch2:
		fmt.Println(v)
	case v:=<- ch3:
		fmt.Println(v)
	}
}
func Getdata( i int,t time.Duration,ch chan int){
	time.Sleep(t*time.Second)
	ch <- i

}
