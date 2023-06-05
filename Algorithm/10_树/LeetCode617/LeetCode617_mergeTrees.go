package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(t1, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.Val += t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}

func main() {
	nodeA1 := new(TreeNode)
	nodeA2 := new(TreeNode)
	nodeA3 := new(TreeNode)
	nodeA4 := new(TreeNode)

	nodeB1 := new(TreeNode)
	nodeB2 := new(TreeNode)
	nodeB3 := new(TreeNode)
	nodeB4 := new(TreeNode)
	nodeB5 := new(TreeNode)

	nodeA1.Val = 1
	nodeA2.Val = 3
	nodeA3.Val = 2
	nodeA4.Val = 5

	nodeB1.Val = 2
	nodeB2.Val = 1
	nodeB3.Val = 3
	nodeB4.Val = 4
	nodeB5.Val = 7

	nodeA1.Left = nodeA2
	nodeA1.Right = nodeA3
	nodeA2.Left = nodeA4

	head := mergeTrees(nodeA1, nodeB1)
	fmt.Println(head)
}
