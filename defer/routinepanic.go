package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	go func() {
		panic("go routine panic")
	}()
	fmt.Println("nothing happened")
}
