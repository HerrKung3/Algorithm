package main

import "fmt"

func findErrorNums(nums []int) []int {
	ans := make([]int, 2)
	cnt := make([]bool, len(nums)+1)
	for _, num := range nums {
		if cnt[num] {
			ans[0] = num
		}
		cnt[num] = true
	}
	for i := 0; i < len(cnt); i++ {
		if cnt[i] == false {
			ans[1] = i
		}
	}
	return ans
}

func main() {
	nums := []int{1, 2, 3, 3, 5}
	ans := findErrorNums(nums)
	fmt.Println(ans)
}
