package main

import (
	"container/heap"
	"fmt"
	"strings"
)

type Nodes [][2]int

func (n *Nodes) Less(i, j int) bool {
	return (*n)[i][1] > (*n)[j][1]
}

func (n *Nodes) Push(x any) {
	*n = append(*n, x.([2]int))
}

func (n *Nodes) Pop() any {
	length := len(*n)
	node := (*n)[length-1]
	*n = (*n)[:length-1]
	return node
}

func (n *Nodes) Len() int {
	return len(*n)
}

func (n *Nodes) Swap(i, j int) {
	(*n)[i], (*n)[j] = (*n)[j], (*n)[i]
}

func frequencySort(s string) string {
	charFreq := make(map[int]int)
	for _, char := range s {
		if _, found := charFreq[int(char)]; found {
			charFreq[int(char)]++
		} else {
			charFreq[int(char)] = 1
		}
	}
	nodes := &Nodes{}
	heap.Init(nodes)
	for char, freq := range charFreq {
		heap.Push(nodes, [2]int{char, freq})
	}
	//fmt.Println(nodes)

	ans := ""
	for nodes.Len() > 0 {
		node := heap.Pop(nodes).([2]int)
		//fmt.Println(node)
		ans += strings.Repeat(string(node[0]), node[1])
	}

	return ans
}

func main() {
	s := "tree"
	ans := frequencySort(s)
	fmt.Println(ans)
}
