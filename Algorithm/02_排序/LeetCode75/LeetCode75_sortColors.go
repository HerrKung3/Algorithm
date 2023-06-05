package main

import "fmt"

func sortColors(nums []int) {
	target := 1
	left, right := -1, len(nums)
	for i := 0; i < right; {
		if nums[i] > target {
			right--
			nums[i], nums[right] = nums[right], nums[i]
		} else if nums[i] < target {
			left++
			nums[i], nums[left] = nums[left], nums[i]
			i++
		} else {
			i++
		}
	}
}

func main() {
	nums := []int{1, 2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}
