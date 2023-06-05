package main

import "fmt"

type MyStack struct {
	Stack []int
}

func (ms *MyStack) Push(node int) {
	ms.Stack = append(ms.Stack, node)
}

func (ms *MyStack) Pop() int {
	n := len(ms.Stack)
	node := ms.Stack[n-1]
	ms.Stack = ms.Stack[:n-1]
	return node
}

func (ms *MyStack) Len() int {
	return len(ms.Stack)
}

func (ms *MyStack) IsEmpty() bool {
	return len(ms.Stack) == 0
}

func (ms *MyStack) Top() int {
	return ms.Stack[len(ms.Stack)-1]
}

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i, _ := range ans {
		ans[i] = -1
	}
	stack := &MyStack{}
	for i := 0; i < n*2-1; i++ {
		num := nums[i%n]
		for !stack.IsEmpty() && num > nums[stack.Top()] {
			preIndex := stack.Pop()
			ans[preIndex] = num
		}
		stack.Push(i % n)
	}
	return ans
}

func main() {
	nums := []int{1, 2, 3, 4, 3}
	ans := nextGreaterElements(nums)
	fmt.Println(ans)
}
