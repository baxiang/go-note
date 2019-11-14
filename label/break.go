package main

import "fmt"

func main() {
	for i:=0;i<2;i++{
		for j:=0;j<3;j++{
			if j>1{
				break
			}
			fmt.Printf("i=%d,j=%d\n",i,j)
		}
	}
}
