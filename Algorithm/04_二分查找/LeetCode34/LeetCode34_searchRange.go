package main

import (
	"fmt"
	"sort"
)

func searchRange(nums []int, target int) []int {
	leftMost := sort.SearchInts(nums, target)
	if leftMost == len(nums) || nums[leftMost] != target {
		return []int{-1, -1}
	}
	rightMost := sort.SearchInts(nums, target+1) - 1
	return []int{leftMost, rightMost}
}

func main() {
	nums := []int{5, 7, 8, 8, 8, 8}
	target := 8
	ans := searchRange(nums, target)
	fmt.Println(ans)
}
