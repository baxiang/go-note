package main

import "fmt"

func binarySearch(nums []int,target int)int{
	if len(nums)==0{
		return -1
	}
	left :=0
	right :=len(nums)-1
	for left<=right{
		mid := left+(right-left)>>1
		if nums[mid]==target{
			return mid
		}else if nums[mid]>target{
			right = mid-1
		}else{
			left = mid +1
		}
	}
	return -1
}


func main() {
	fmt.Println(binarySearch([]int{1,3,4,5,6,7},4))
	fmt.Println(binarySearch([]int{1,3,4,5,6,7},8))
}
