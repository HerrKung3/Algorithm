package main

import "fmt"

func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for slow, fast = nums[slow], nums[nums[fast]]; slow != fast; slow, fast = nums[slow], nums[nums[fast]] {
	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

func main() {
	nums := []int{3, 1, 3, 4, 2}
	ans := findDuplicate(nums)
	fmt.Println(ans)
}
