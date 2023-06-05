package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

func printListNode(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Println(cur)
		cur = cur.Next
	}
}

func main() {
	node1 := new(ListNode)
	node2 := new(ListNode)
	node3 := new(ListNode)
	node4 := new(ListNode)
	node5 := new(ListNode)

	node1.Val = 1
	node2.Val = 1
	node3.Val = 2
	node4.Val = 3
	node5.Val = 3

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	ans := deleteDuplicates(node1)
	printListNode(ans)
}
