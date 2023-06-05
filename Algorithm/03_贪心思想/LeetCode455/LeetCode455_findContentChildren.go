package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	ans := 0
	sort.Ints(g)
	sort.Ints(s)

	for i, j := 0, 0; i < len(g) && j < len(s); {
		if s[j] >= g[i] {
			ans++
			i++
			j++
		} else {
			j++
		}
	}
	return ans
}

func main() {
	g := []int{3, 6, 5, 7, 2, 1}
	s := []int{2, 1, 6, 3, 1, 2, 2, 4}
	ans := findContentChildren(g, s)
	fmt.Println(ans)
}
