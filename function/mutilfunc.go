package main

import "fmt"

func min(a ...int) int {
	if len(a) == 0 {
		return 0

	}
	min := a[0]
	for _, v := range a {
		if min > v {
			min = v
		}
	}
	return min
}
func main() {
	n := min(17, 12, 5, 9, 8)
	fmt.Println(n)
	arr := []int{26, 17, 35, 6, 18}
	r := min(arr...)
	fmt.Println(r)
	m := min(arr[:3]...)
	fmt.Println(m)
}
