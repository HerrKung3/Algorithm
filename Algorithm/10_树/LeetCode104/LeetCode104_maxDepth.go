package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	node1 := new(TreeNode)
	node2 := new(TreeNode)
	node3 := new(TreeNode)
	node4 := new(TreeNode)
	node5 := new(TreeNode)

	node1.Left = node2
	node1.Right = node3
	node3.Left = node4
	node3.Right = node5

	node1.Val = 1
	node2.Val = 2
	node3.Val = 3
	node4.Val = 4
	node5.Val = 5

	ans := maxDepth(node1)
	fmt.Println(ans)
}
