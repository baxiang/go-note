package main

import "fmt"

func main() {
	//done := make(chan int)

	go func(){
		fmt.Println("你好, 世界")
		//done<-1
	}()

	//<-done
}

