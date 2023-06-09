package main

import "fmt"

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	dp := make([]int, n)
	dp[0], dp[1] = nums[0], max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{2, 7, 9, 3, 1, 3}
	ans := rob(nums)
	fmt.Println(ans)
}
