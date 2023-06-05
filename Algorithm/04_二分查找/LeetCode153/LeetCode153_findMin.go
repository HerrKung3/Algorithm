package main

import "fmt"

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (right-left)/2 + left
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}

func main() {
	nums := []int{4, 5, 6, 1, 2, 3}
	ans := findMin(nums)
	fmt.Println(ans)
}
