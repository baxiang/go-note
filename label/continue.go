package main

import (
	"fmt"
)

func main() {
	FirstNames := []string{"张", "李", "王"}
	LastNames := []string{"三", "四", "五"}
	Loop:
	for i, firstName := range FirstNames {
		for j, lastName := range LastNames {
			if i == 1 && j==1 {
				continue Loop
			}
			fmt.Printf("%d-%d:%s%s\n",i,j, firstName, lastName)
		}
	}
}
