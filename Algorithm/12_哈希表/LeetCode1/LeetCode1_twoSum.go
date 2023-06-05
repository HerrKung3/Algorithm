package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int, 0)
	var ans []int
	for i := 0; i < len(nums); i++ {
		sum := target - nums[i]
		if index, found := hashMap[sum]; found {
			ans = append(ans, index, i)
		}
		hashMap[nums[i]] = i
	}
	return ans
}

func main() {
	nums := []int{2, 11, 7, 15}
	ans := twoSum(nums, 9)
	fmt.Println(ans)
}
