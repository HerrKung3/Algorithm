package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	ans := 0
	for i, _ := range nums {
		count := 0
		for i < len(nums) && nums[i] == 1 {
			count++
			i++
		}
		ans = Max(ans, count)
	}
	return ans
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 1, 0, 1, 0, 1, 1, 1, 1, 0}
	ans := findMaxConsecutiveOnes(nums)
	fmt.Println(ans)
}
