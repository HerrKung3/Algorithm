package main

import (
	"container/heap"
	"fmt"
)

type Nodes [][2]int

func topKFrequent(nums []int, k int) []int {
	numsFreq := make(map[int]int, 0)
	for _, n := range nums {
		if _, found := numsFreq[n]; found {
			numsFreq[n]++
		} else {
			numsFreq[n] = 1
		}
	}
	//fmt.Println(numsFreq)
	return MaxHeap(numsFreq, k)
}

func MaxHeap(numFreq map[int]int, k int) (ans []int) {
	nodes := &Nodes{}
	heap.Init(nodes)
	for n, freq := range numFreq {
		heap.Push(nodes, [2]int{n, freq})
	}
	//fmt.Println(nodes)
	for i := 0; i < k; i++ {
		num := heap.Pop(nodes)
		//fmt.Println(num)
		ans = append(ans, num.([2]int)[0])
	}
	return
}

func (n *Nodes) Len() int {
	return len(*n)
}

func (n *Nodes) Less(i, j int) bool {
	return (*n)[i][1] > (*n)[j][1]
}

func (n *Nodes) Swap(i, j int) {
	(*n)[i], (*n)[j] = (*n)[j], (*n)[i]
}

func (n *Nodes) Push(node interface{}) {
	*n = append(*n, node.([2]int))
}

func (n *Nodes) Pop() interface{} {
	old := *n
	length := len(old)
	node := old[length-1]
	*n = old[:length-1]
	return node
}

func main() {
	nums := []int{3, 0, 1, 0}
	ans := topKFrequent(nums, 1)
	fmt.Println(ans)
}
