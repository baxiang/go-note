package main

import (
	"fmt"
)

type Node struct {
	data int
	node *Node
}

func StackRevert(head *Node) *Node{
	h := head
	currNode := head
	for {
		tmp :=currNode.node
		if tmp==nil {
			break
		}
		currNode.node= currNode
		currNode = tmp
	}
	h.node= nil
	return currNode
}

func main() {
	 four := &Node{data:4,node:nil}
	 three :=&Node{data:3,node:four}
	 two :=&Node{data:2,node:three}
	 one :=&Node{data:1,node:two}
	// head := one
	head := StackRevert(one)
	for head!=nil{
		fmt.Println(head.data)
		head = head.node
	}
}
