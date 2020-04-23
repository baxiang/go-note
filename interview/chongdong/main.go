package main

import (
	"fmt"
	"math/rand"
	"time"
)

func produce(c chan int){
	for{
		n:= rand.Intn(1000)
		c<-n
		i++
		time.Sleep(time.Second)
	}
}

func consume(c chan int){

	for{
		select {
		case k:=<-c:
			j++
			time.Sleep(time.Duration(k)*time.Millisecond)
			//fmt.Println(v)
		}

	}
}
func watch(){
	for {
		fmt.Println("producer num", i)
		fmt.Println("consumer num", j)
		time.Sleep(time.Second*5)
	}
}
var i int
var j int
func main() {
	rand.Seed(time.Now().UnixNano())
	ch := make(chan int,10)
	for i:=0;i<100;i++{
		go produce(ch)
	}
	go consume(ch)
	watch()
	select {
	}

}
