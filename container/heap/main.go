package main

import (
	"container/heap"
	"fmt"
)

type myHeap []int
func(h myHeap)Len()int{
	return len(h)
}
func(h myHeap)Swap(i,j int){
	h[i],h[j] = h[j],h[i]
}
func(h myHeap)Less(i,j int)bool{
	return h[i]>h[j]
}
func(h *myHeap)Push(v interface{}){
	*h = append(*h,v.(int))
}
func(h *myHeap)Pop()(v interface{}){
	*h,v =(*h)[:len(*h)-1],(*h)[len(*h)-1]
	return
}

// 按层来遍历和打印堆数据，第一行只有一个元素，即堆顶元素
func (h myHeap) printHeap() {
	n := 1
	levelCount := 1
	for n <= h.Len() {
		fmt.Println(h[n-1 : n-1+levelCount])
		n += levelCount
		levelCount *= 2
	}
}
// 大顶堆小堆的结果主要swap
func main() {
	list :=[7]int{13,12,45,24,11,9,20}
	hq := make(myHeap,0)
	for i :=range list{
		hq.Push(list[i])
	}
	heap.Init(&hq)
	hq.printHeap()
}
