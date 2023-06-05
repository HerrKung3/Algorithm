package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type Stack struct {
	Queue []int
}

func (s *Stack) Pop() int {
	node := s.Queue[0]
	s.Queue = s.Queue[1:]
	return node
}

func (s *Stack) Push(n int) {
	s.Queue = append(s.Queue, n)
}

func (s *Stack) Len() int {
	return len(s.Queue)
}

func (s *Stack) IsEmpty() bool {
	return len(s.Queue) == 0
}

func canReach(arr []int, start int) bool {
	path := make([]bool, len(arr))
	queue := Stack{}
	queue.Push(start)
	path[start] = true
	for !queue.IsEmpty() {
		size := queue.Len()
		cur := queue.Pop()
		for size > 0 {
			size--
			if arr[cur] == 0 {
				return true
			} else {
				if cur+arr[cur] < len(arr) && cur+arr[cur] >= 0 && !path[cur+arr[cur]] {
					path[cur+arr[cur]] = true
					queue.Push(cur + arr[cur])
				}
				if cur-arr[cur] < len(arr) && cur-arr[cur] >= 0 && !path[cur-arr[cur]] {
					path[cur-arr[cur]] = true
					queue.Push(cur - arr[cur])
				}
			}
		}
	}
	return false
}

func main() {
	//arr := []int{4, 2, 3, 0, 3, 1, 2}
	//start := 0
	//arr := []int{4, 2, 3, 0, 3, 1, 2}
	//start := 5
	arr := []int{3, 0, 2, 1, 2}
	start := 2
	ans := canReach(arr, start)
	fmt.Println(ans)
}
