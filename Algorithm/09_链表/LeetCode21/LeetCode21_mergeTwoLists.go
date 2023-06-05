package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(head1 *ListNode, head2 *ListNode) *ListNode {
	if head1 == nil {
		return head2
	} else if head2 == nil {
		return head1
	}

	var head, tail *ListNode
	if head1.Val > head2.Val {
		head, tail, head2 = head2, head2, head2.Next
	} else {
		head, tail, head1 = head1, head1, head1.Next
	}
	for head1 != nil && head2 != nil {
		if head1.Val > head2.Val {
			tail.Next, head2 = head2, head2.Next
		} else {
			tail.Next, head1 = head1, head1.Next
		}
		tail = tail.Next
	}
	if head1 != nil {
		tail.Next = head1
	}
	if head2 != nil {
		tail.Next = head2
	}
	//printList(head)
	return head
}

//func printList(head *ListNode) {
//	cur := head
//	for cur != nil {
//		fmt.Println(cur.Val)
//		cur = cur.Next
//	}
//}

func main() {
	nodeA1 := new(ListNode)
	nodeA2 := new(ListNode)
	nodeA3 := new(ListNode)

	nodeB1 := new(ListNode)
	nodeB2 := new(ListNode)
	nodeB3 := new(ListNode)

	nodeA1.Val = 1
	nodeA2.Val = 2
	nodeA3.Val = 5

	nodeB1.Val = 1
	nodeB2.Val = 3
	nodeB3.Val = 4

	nodeA1.Next = nodeA2
	nodeA2.Next = nodeA3

	nodeB1.Next = nodeB2
	nodeB2.Next = nodeB3

	ans := mergeTwoLists(nodeA1, nodeB1)
	fmt.Println(ans)
}
