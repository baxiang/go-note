package main

import (
	"fmt"
	"sort"
)

func searchTopK(nums []int,k int)int {
	//m :=make(map[int]struct{})
	//var list []int
	//for _,v:=range nums{
	//	if _,ok:=m[v];!ok{
	//		list = append(list,v)
	//		m[v] = struct{}{}
	//	}
	//}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i]>nums[j]
	})
	//fmt.Println(nums)
	return nums[k-1]
}

func main() {
	fmt.Println(searchTopK([]int{3,2,1,5,6,4},2))
	fmt.Println(searchTopK([]int{3,2,3,1,2,4,5,5,6},4))
}
