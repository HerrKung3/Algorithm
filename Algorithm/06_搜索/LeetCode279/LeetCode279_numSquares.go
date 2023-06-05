package main

import "fmt"

type Queue struct {
	Nodes []interface{}
}

func numSquares(n int) int {
	squares := generateSquares(n)
	queue := Queue{}
	path := make([]bool, n+1)

	queue.Push(n)
	path[n] = true

	level := 0
	for !queue.IsEmpty() {
		level++
		size := queue.Len()
		for size > 0 {
			size--
			curNode := queue.Pop()
			for _, s := range squares {
				next := curNode.(int) - s
				if next < 0 {
					break
				}
				if next == 0 {
					return level
				}
				if path[next] {
					continue
				}
				path[next] = true
				queue.Push(next)
			}
		}
	}
	return level
}

// 生成n以内的整数平方 1 4 9...
func generateSquares(n int) []int {
	squares := make([]int, 0)
	square := 1
	diff := 3
	for square <= n {
		squares = append(squares, square)
		square += diff
		diff += 2
	}
	return squares
}

func (arr *Queue) Push(node interface{}) {
	arr.Nodes = append(arr.Nodes, node)
}

func (arr *Queue) Pop() interface{} {
	first := arr.Nodes[0]
	arr.Nodes = arr.Nodes[1:]
	return first
}

func (arr *Queue) Len() int {
	return len(arr.Nodes)
}

func (arr *Queue) IsEmpty() bool {
	return len(arr.Nodes) == 0
}

func main() {
	ans := numSquares(12)
	fmt.Println(ans)
}
