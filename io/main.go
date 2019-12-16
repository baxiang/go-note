package main

func main() {
	c := make(chan int,0)
	<-c
}
