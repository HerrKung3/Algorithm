package main

import (
	"fmt"
	"math/rand"
)

// 快速排序
func findKthLargest2(nums []int, k int) int {
	quickSort(nums, 0, len(nums)-1)
	return nums[len(nums)-k]
}

func quickSort(nums []int, start, end int) {
	if start < end {
		randomIndex := rand.Intn(end-start+1) + start
		nums[randomIndex], nums[end] = nums[end], nums[randomIndex]
		left, right := partition(nums, start, end)
		quickSort(nums, start, left)
		quickSort(nums, right+1, end)
	}
}

func partition(nums []int, start, end int) (int, int) {
	target := nums[end]
	less, more := start-1, end
	for n := start; n <= more; {
		if nums[n] > target {
			nums[n], nums[more] = nums[more], nums[n]
			more--
		} else if nums[n] < target {
			less++
			nums[n], nums[less] = nums[less], nums[n]
			n++
		} else {
			n++
		}
	}
	return less, more
}

func main() {
	//nums := []int{3, 2, 3, 1, 2, 4, 5, 7, 6}
	nums := []int{3, 2, 1, 5, 6, 4}
	ans := findKthLargest2(nums, 2)
	fmt.Println(ans)
}
