package main

import (
	"fmt"
	"sort"
)

func reconstructQueue(people [][]int) (ans [][]int) {
	sort.Slice(people, func(i, j int) bool {
		x, y := people[i], people[j]
		return x[0] > y[0] || (x[0] == y[0] && x[1] < y[1])
	})
	//fmt.Println(people)
	for _, person := range people {
		idx := person[1]
		//golang中插入的实现
		ans = append(ans[:idx], append([][]int{person}, ans[idx:]...)...)
	}
	return ans
}

func main() {
	people := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	ans := reconstructQueue(people)
	fmt.Println(ans)
}
