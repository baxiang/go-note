package main

import "fmt"

func main() {
Loop:
	for i:=0;i<3 ;i++  {
		for j :=0;j<=i;j++{
			if j==2 {
				break Loop
			}
			fmt.Printf("i=%d,j=%d\n",i,j)
		}
	}

}
