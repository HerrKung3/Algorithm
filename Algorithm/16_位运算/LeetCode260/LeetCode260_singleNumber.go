package main

import "fmt"

func singleNumber(nums []int) []int {
	xorSum := 0
	for _, num := range nums {
		xorSum ^= num
	}
	lsb := xorSum & -xorSum
	type1, type2 := 0, 0
	for _, num := range nums {
		if lsb&num > 0 {
			type1 ^= num
		} else {
			type2 ^= num
		}
	}
	return []int{type1, type2}
}

func main() {
	nums := []int{1, 2, 1, 3, 2, 5}
	ans := singleNumber(nums)
	fmt.Println(ans)
}
