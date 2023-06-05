package main

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)
	//left[i]表示nums[i]之前元素乘积
	left := make([]int, len(nums))
	//right[i]表示nums[i]之后元素乘积
	right := make([]int, len(nums))
	ans := make([]int, len(nums))
	left[0], right[n-1] = 1, 1
	for i := 1; i < len(left); i++ {
		left[i] = left[i-1] * nums[i-1]
	}
	for i := n - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}
	for i := 0; i < n; i++ {
		ans[i] = left[i] * right[i]
	}
	return ans
}

func main() {
	nums := []int{-1, 1, 0, -3, 3}
	ans := productExceptSelf(nums)
	fmt.Println(ans)
}
