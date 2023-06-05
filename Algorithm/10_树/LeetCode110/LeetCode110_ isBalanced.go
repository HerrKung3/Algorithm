package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方法一：自顶向下的递归
// 时间复杂度：O(n^2) 空间复杂度：O(n)
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return abs(height(root.Left)-height(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(height(root.Left), height(root.Right)) + 1
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

	node1.Val = 1
	node2.Val = 2
	node3.Val = 3
	node4.Val = 4
	node5.Val = 5

	node1.Left = node2
	node1.Right = node3
	node3.Left = node4
	node3.Right = node5

	ans := isBalanced(node1)
	fmt.Println(ans)
}
