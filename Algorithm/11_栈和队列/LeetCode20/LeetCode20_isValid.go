package main

import "fmt"

type MyStack struct {
	Stack []string
}

func (ms *MyStack) Pop() string {
	n := len(ms.Stack)
	node := ms.Stack[n-1]
	ms.Stack = ms.Stack[:n-1]
	return node
}

func (ms *MyStack) Push(node string) {
	ms.Stack = append(ms.Stack, node)
}

func (ms *MyStack) Len() int {
	return len(ms.Stack)
}

func (ms *MyStack) IsEmpty() bool {
	return len(ms.Stack) == 0
}

func (ms *MyStack) Top() string {
	return ms.Stack[len(ms.Stack)-1]
}

func isValid(s string) bool {
	stack := &MyStack{}
	for i := 0; i < len(s); i++ {
		switch char := string(s[i]); {
		case char == "(" || char == "[" || char == "{":
			stack.Push(char)
		case char == ")":
			if stack.IsEmpty() || stack.Top() != "(" {
				return false
			}
			stack.Pop()
		case char == "]":
			if stack.IsEmpty() || stack.Top() != "[" {
				return false
			}
			stack.Pop()
		case char == "}":
			if stack.IsEmpty() || stack.Top() != "{" {
				return false
			}
			stack.Pop()
		}
	}
	return stack.IsEmpty()
}

func main() {
	s := "({}[]([]))"
	ans := isValid(s)
	fmt.Println(ans)
}
