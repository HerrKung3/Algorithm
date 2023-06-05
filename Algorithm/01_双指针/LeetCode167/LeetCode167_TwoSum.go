package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	var ans []int
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum < target {
			left++
		} else if sum > target {
			right--
		} else {
			ans = append(ans, left+1, right+1)
			break
		}
	}
	return ans
}

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 9
	ans := twoSum(numbers, target)
	fmt.Println(ans)
}
