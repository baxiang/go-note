package main

import "fmt"
type people struct {
	Id int
}


func main() {

	n :=[]int{1,2,3}
	//n = append(n[:2],n[3:]...)
	//fmt.Println(n)
	l :=make([]people,0)
	for _,v:=range n{
		l = append(l,people{Id:v})
	}
	var i =1
	var idx =0
	for len(l)>0 {
		if len(l) == 1 {
			fmt.Println(l[0].Id)
		}
		if i == 3 {
			l = append(l[:idx], l[idx+1:]...)
			i = 1
			idx--
		}
		idx++
		i++
	}
}
