package main

import "fmt"

func index(a []int, t int)int {
	 left := 0
	 right := len(a)-1
	for left<=right  {
		mid :=left+(right-left)/2
		if a[mid]==t {
			return mid
		}else if a[mid]>t {
			right = mid-1
		}else {
			left = mid+1
		}
	}
	return left
}

func main() {
   fmt.Println(index([]int{1,7,17,19},10))
}
