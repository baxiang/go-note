package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	for i:=1;i<5;i++{
		l.PushBack(i)
	}
	fmt.Println(l.Front().Value)//输出首部元素的值
	fmt.Println(l.Back().Value)//输出尾部元素的值
	// 遍历链表
	for e :=l.Front();e!=nil;e =e.Next(){
		fmt.Print(e.Value)
	}
	fmt.Println("")
	l.InsertBefore(0, l.Front())  //首部元素之前插入0
	l.MoveToBack(l.Front()) //将0元素移动到末尾
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
	}
	fmt.Println("")

	l2 := list.New()
	l2.PushBackList(l) //将l中元素放在l2的末尾
	l2.MoveToFront(l2.Back()) // 将末尾元素移到首部
	for e := l2.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
	}
	fmt.Println("")
	l.Init() //清空l
	fmt.Println(l.Len()) //0
}
