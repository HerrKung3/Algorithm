package main

import "fmt"

func singleNonDuplicate(nums []int) int {
	n := 0
	for i, _ := range nums {
		n = n ^ nums[i]
	}
	return n
}

func main() {
	nums := []int{1, 1, 2, 2, 3, 3, 4, 4, 8}
	ans := singleNonDuplicate(nums)
	fmt.Println(ans)
}
