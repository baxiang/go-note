package main

import "fmt"

type TreeNode struct {
	 Val int
	 Left *TreeNode
	 Right *TreeNode
	 isLeft bool
}


func PrintLeftNode(root *TreeNode){
	 stack :=make([]*TreeNode,0)
	 for root!= nil||len(stack)>0{
	 	 for root!= nil{
			 stack = append(stack,root)
			 if root.Left!= nil{
				 root = root.Left
				 root.isLeft = true
			 }
		 }
		 node := stack[len(stack)-1]
		 if node.isLeft{
		 	fmt.Println(node)
		 }
		 stack = stack[:len(stack)-1]
		 root = root.Left
	 }
}


func main() {
	
}
