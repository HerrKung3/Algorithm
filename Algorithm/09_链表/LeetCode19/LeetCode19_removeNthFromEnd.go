package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	height := 0
	cur := head
	for cur != nil {
		height++
		cur = cur.Next
	}
	if height < n {
		return nil
	}
	//给头节点创造一个前节点, 从头节点遍历，避免当删除的是倒数一个时，cur2.Next = cur2.Next.Next引发错误
	dummy := &ListNode{0, head}
	cur = dummy
	count := height - n
	for count > 0 {
		cur = cur.Next
		count--
	}
	cur.Next = cur.Next.Next
	return dummy.Next
}

func printList(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Println("current Node = ", cur)
		cur = cur.Next
	}
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

	head := removeNthFromEnd(node1, 6)
	printList(head)
}
