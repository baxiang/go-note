package main


import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func add(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = revertNode(l1)
	l2 = revertNode(l2)
	dummyHead := &ListNode{Val: -1,}
	var carry int
	for l1 != nil || l2 != nil {
		var a int
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}
		var b int
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}
		sum := a + b + carry
		carry = sum / 10
		node := &ListNode{
			Val: sum % 10,
		}
		head := dummyHead.Next
		node.Next = head
		dummyHead.Next = node
	}
	if carry > 0 {
		node := &ListNode{
			Val: carry,
		}
		head := dummyHead.Next
		node.Next = head
		dummyHead.Next = node
	}
	return dummyHead.Next
}

func revertNode(h *ListNode) *ListNode {
	var pre *ListNode
	curr :=h
	for curr!=nil{
		pre,curr,curr.Next = curr,curr.Next,pre
	}
	return pre
}

func defer_call(){
	val :="111"
	defer func() {
		fmt.Println(3)
		fmt.Printf("%p\n",&val)
		fmt.Println(val)
	}()
	defer fmt.Printf("%p\n",&val)
	val ="222"
	defer func(val string) {
		fmt.Println(2)
		fmt.Printf("%p\n",&val)
		fmt.Println(val)
	}(val)
	val ="333"
	defer func() {
		fmt.Printf("%p\n",&val)
		fmt.Println(val)
	}()
}

func main() {
	//a1 :=  &ListNode{
	//	Val:  8,
	//}
	//b1 :=  &ListNode{
	//	Val:  0,
	//}
	//a1.Next = b1
	//c1 :=  &ListNode{
	//	Val:  6,
	//}
	//b1.Next = c1
	//
	//
	//a2 :=  &ListNode{
	//	Val:  1,
	//}
	//b2 :=  &ListNode{
	//	Val:  2,
	//}
	//a2.Next = b2
	//c2 :=  &ListNode{
	//	Val:  3,
	//}
	//b2.Next = c2
	//
	//r := add(a1,a2)
	//for r != nil{
	//	fmt.Println(r.Val)
	//	r = r.Next
	//}
	defer_call()
	//var c chan int
	//<-c
}
