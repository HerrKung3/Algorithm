package main

import (
	"container/heap"
	"fmt"
)

type IHeap []int

func (h *IHeap) Len() int {
	return len(*h)
}

func (h *IHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *IHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IHeap) Push(node interface{}) {
	*h = append(*h, node.(int))
}

func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	last := old[n-1]
	*h = old[:n-1]
	return last
}

func kthSmallest(matrix [][]int, k int) int {
	M := len(matrix)
	N := len(matrix[0])
	h := &IHeap{}
	heap.Init(h)
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			heap.Push(h, matrix[i][j])
		}
	}
	//fmt.Println("heap=", h)
	for n := 0; n < M*N-k; n++ {
		_ = heap.Pop(h)
	}
	return (*h)[0]
}

func main() {
	matrix := [][]int{
		{1, 5, 9},
		{10, 11, 13},
		{12, 13, 15}}
	ans := kthSmallest(matrix, 5)
	fmt.Println(ans)
}
