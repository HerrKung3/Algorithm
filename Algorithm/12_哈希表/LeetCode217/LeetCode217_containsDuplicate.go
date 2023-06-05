package main

import "fmt"

func containsDuplicate(nums []int) bool {
	numFreq := make(map[int]int, 0)
	for _, n := range nums {
		if _, found := numFreq[n]; found {
			return true
		} else {
			numFreq[n] = 1
		}
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 4}
	ans := containsDuplicate(nums)
	fmt.Println(ans)
}
