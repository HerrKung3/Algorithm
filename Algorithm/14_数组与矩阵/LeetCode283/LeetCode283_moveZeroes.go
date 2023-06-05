package main

import "fmt"

func moveZeroes(nums []int) {
	for i, j := 0, 1; i < len(nums) && j < len(nums); j++ {
		if nums[i] != 0 {
			i++
		}
		if nums[j] != 0 && nums[i] == 0 {
			nums[j], nums[i] = nums[i], nums[j]
			i++
		}
	}
}

func main() {
	nums := []int{0, 1, 2, 3, 1, 0}
	moveZeroes(nums)
	fmt.Println(nums)
}
