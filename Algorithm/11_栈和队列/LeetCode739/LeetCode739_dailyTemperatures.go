package main

import "fmt"

type MyStack struct {
	Stack []int
}

func (ms *MyStack) Top() int {
	return ms.Stack[len(ms.Stack)-1]
}

func (ms *MyStack) Push(node int) {
	ms.Stack = append(ms.Stack, node)
}

func (ms *MyStack) Pop() int {
	n := len(ms.Stack)
	last := ms.Stack[n-1]
	ms.Stack = ms.Stack[:n-1]
	return last
}

func (ms *MyStack) IsEmpty() bool {
	return len(ms.Stack) == 0
}

func (ms *MyStack) Len() int {
	return len(ms.Stack)
}

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)
	stack := &MyStack{}
	for i := 0; i < n; i++ {
		t := temperatures[i]
		for !stack.IsEmpty() && t > temperatures[stack.Top()] {
			preIndex := stack.Pop()
			ans[preIndex] = i - preIndex
		}
		stack.Push(i)
	}
	return ans
}

func main() {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	ans := dailyTemperatures(temperatures)
	fmt.Println(ans)
}
