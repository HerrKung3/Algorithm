package main

import (
	"fmt"
)

func Solution(nums []int) []int {
	ans := make([]int, 500)
	for _, num := range nums {
		if ans[num] == 0 {
			ans[num] = 1
		}
	}
	return ans
}

func main() {
	a := 0
	n := 0
	_, _ = fmt.Scan(&n)
	nums := make([]int, 0)
	var ans []int
	for n > 0 {
		_, _ = fmt.Scan(&a)
		nums = append(nums, a)
		n--
		//fmt.Printf("%d\n", a + b)
	}
	ans = Solution(nums)
	for i, _ := range ans {
		if ans[i] != 0 {
			fmt.Printf("%d\n", i)
		}
	}
	//fmt.Println("nums=", nums)
}
