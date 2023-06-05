package main

import "fmt"

func arrayNesting(nums []int) int {
	Nesting := make([]bool, len(nums))
	ans := 0

	for i, _ := range nums {
		cnt := 0
		for !Nesting[i] {
			cnt++
			Nesting[i] = true
			i = nums[i]
		}
		ans = max(ans, cnt)
		cnt = 0
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{5, 4, 0, 3, 1, 6, 2}
	ans := arrayNesting(nums)
	fmt.Println(ans)
}
