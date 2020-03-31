package main

import "fmt"

type A struct {
	bList []*B
}

type B struct {
	word string
}

func main() {
	bList  := []*B{&B{word:"1"},&B{word:"2"}}
	a :=&A{bList:nil}
	a.bList = append(a.bList,bList...)
	m :=make(map[string]*A)
	m["A"]= a
	var c []*A
	currA := m["A"]
	aa :=&A{bList:currA.bList}
	aa.bList = append(aa.bList,&B{word:"3"})
	c = append(c, aa)
	currB := m["A"]
	c = append(c, currB)
	for _,v :=range c{
		fmt.Println(len(v.bList))
	}
}
