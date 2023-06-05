package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func main() {
	node1 := new(ListNode)
	node2 := new(ListNode)
	node3 := new(ListNode)
	node4 := new(ListNode)
	node5 := new(ListNode)

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	node1.Val = 1
	node2.Val = 2
	node3.Val = 3
	node4.Val = 4
	node5.Val = 5

	ans := reverseList(node1)
	fmt.Println(ans)
}
