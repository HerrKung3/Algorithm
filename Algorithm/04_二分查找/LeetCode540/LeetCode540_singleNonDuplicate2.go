package main

import "fmt"

func singleNonDuplicate2(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == nums[mid^1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}

func main() {
	nums := []int{1, 1, 2, 3, 3, 4, 4, 8, 8}
	ans := singleNonDuplicate2(nums)
	fmt.Println(ans)
}
