package main

import (
	"fmt"
)
func quickSort(a []int)[]int{
	if len(a)<=1{
		return a
	}
	pro := a[0]
	mid :=[]int{}
	left :=[]int{}
	right :=[]int{}
	for i:=0;i<len(a);i++{
		if a[i]==pro{
			mid = append(mid,a[i])
		}else if a[i]<pro{
			right = append(right,a[i])
		}else{
			left = append(left,a[i])
		}
	}
	left = quickSort(left)
	left = append(left,mid...)
	right = quickSort(right)
	return append(left,right...)
}

func searchK(a []int,k int)int{
	a = quickSort(a)
	return a[k-1]
}

func main() {
	fmt.Println(searchK([]int{4,2,3},1))
}