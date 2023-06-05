package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}

func main() {
	nodeA1 := new(ListNode)
	nodeA2 := new(ListNode)
	nodeA3 := new(ListNode)
	nodeA4 := new(ListNode)

	nodeB1 := new(ListNode)

	nodeA1.Val = 1
	nodeA2.Val = 2
	nodeA3.Val = 3
	nodeA4.Val = 4
	nodeB1.Val = 5

	nodeA1.Next = nodeA2
	nodeA2.Next = nodeA3
	nodeA3.Next = nodeA4
	nodeB1.Next = nodeA3

	ans := getIntersectionNode(nodeA1, nodeB1)
	fmt.Println(ans)
}
