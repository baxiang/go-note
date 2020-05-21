package main

import "fmt"

type TreeNode struct {
	 V int
	 Left *TreeNode
	 Right *TreeNode
}

func traverse(root *TreeNode){
	if root==nil {
		return
	}
	fmt.Println(root.V)
	traverse(root.Left)
	traverse(root.Right)
}
func traverse1(root *TreeNode){
	if root==nil {
		return
	}
	stack :=make([]*TreeNode,0)
	for root!= nil||len(stack)>0{
		if root!=nil{
			stack = append(stack,root)
			root = root.Left
		}else {
			node := stack[len(stack)-1]
			stack =stack[:len(stack)]
			fmt.Println(node.V)
			root = root.Right
		}
	}
}

func main() {
	
}
