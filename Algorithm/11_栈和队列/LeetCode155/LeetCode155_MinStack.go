package main

import (
	"fmt"
	"math"
)

type MinStack struct {
	Stack []int
	Min   []int
}

func Constructor() MinStack {
	return MinStack{Stack: []int{}, Min: []int{math.MaxInt64}}
}

func (ms *MinStack) Push(val int) {
	n := len(ms.Min)
	ms.Stack = append(ms.Stack, val)
	top := ms.Min[n-1]
	ms.Min = append(ms.Min, min(top, val))
}

func (ms *MinStack) Pop() {
	ms.Stack = ms.Stack[:len(ms.Stack)-1]
	ms.Min = ms.Min[:len(ms.Min)-1]
}

func (ms *MinStack) Top() int {
	return ms.Stack[len(ms.Stack)-1]
}

func (ms *MinStack) GetMin() int {
	return ms.Min[len(ms.Min)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	minstack := Constructor()
	minstack.Push(-2)
	minstack.Push(0)
	minstack.Push(-3)
	min1 := minstack.GetMin()
	fmt.Println(min1)
	minstack.Pop()
	top1 := minstack.Top()
	fmt.Println(top1)
	min2 := minstack.GetMin()
	fmt.Println(min2)
}
