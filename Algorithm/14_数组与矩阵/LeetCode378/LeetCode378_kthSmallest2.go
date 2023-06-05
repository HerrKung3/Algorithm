// 数组 从二维数组中找出第k小元素
// 每一行都是有序数组，从n行中归并排序到第k个元素，使用小根堆维护，优化时间复杂度
package main

import (
	"container/heap"
	"fmt"
)

func kthSmallest2(matrix [][]int, k int) int {
	h := &Nodes{}
	for i := 0; i < len(matrix); i++ {
		heap.Push(h, [3]int{matrix[i][0], i, 0})
	}

	for i := 0; i < k-1; i++ {
		now := heap.Pop(h).([3]int)
		if now[2] != len(matrix)-1 {
			heap.Push(h, [3]int{matrix[now[1]][now[2]+1], now[1], now[2] + 1})
		}
	}
	return heap.Pop(h).([3]int)[0]
}

type Nodes [][3]int

func (h Nodes) Len() int           { return len(h) }
func (h Nodes) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h Nodes) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Nodes) Push(x interface{}) {
	*h = append(*h, x.([3]int))
}

func (h *Nodes) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	matrix := [][]int{
		{1, 3, 7, 11},
		{2, 6, 8, 13},
		{9, 12, 14, 15},
		{11, 16, 17, 19},
	}
	res := kthSmallest2(matrix, 7)
	fmt.Println(res)
}
