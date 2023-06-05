package main

import "fmt"

func findLHS(nums []int) int {
	hashMap := map[int]int{}

	LHS := 0
	for _, num := range nums {
		if _, found := hashMap[num]; found {
			hashMap[num]++
		} else {
			hashMap[num] = 1
		}
		if _, found := hashMap[num-1]; found {
			LHS = max(LHS, hashMap[num]+hashMap[num-1])
		}
		if _, found := hashMap[num+1]; found {
			LHS = max(LHS, hashMap[num]+hashMap[num+1])
		}
	}
	return LHS
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 3, 2, 2, 5, 2, 3, 7}
	ans := findLHS(nums)
	fmt.Println(ans)
}
