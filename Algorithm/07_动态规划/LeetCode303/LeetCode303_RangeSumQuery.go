package main

import "fmt"

type NumArray struct {
	nums []int
	dp   []int
}

func Constructor(nums []int) NumArray {
	numArray := NumArray{
		nums: nums,
		dp:   make([]int, len(nums)+1),
	}
	numArray.dp[0] = 0
	for i := 1; i <= len(numArray.nums); i++ {
		numArray.dp[i] += nums[i-1] + numArray.dp[i-1]
	}
	return numArray
}

func (n *NumArray) SumRange(left int, right int) int {
	return n.dp[right+1] - n.dp[left]
}

func main() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	numArray := Constructor(nums)
	sum1 := numArray.SumRange(2, 5)
	fmt.Println(sum1)
}
