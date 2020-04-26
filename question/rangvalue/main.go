package main

import "fmt"

type User struct {
	Id int
	Name string
}
func main() {
	stus :=[]User{
		User{Id:1,Name:"zhang"},
		User{Id:2,Name:"li"},
		User{Id:3,Name:"wang"},
	}
	for _,stu :=range stus{
		stu.Id = stu.Id+1
	}
	for _,stu :=range stus{
		fmt.Println(stu)
	}
}
