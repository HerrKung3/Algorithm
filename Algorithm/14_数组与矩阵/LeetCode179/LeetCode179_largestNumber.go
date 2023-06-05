package main

import (
	"fmt"
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		sx, sy := 10, 10
		for sx <= x {
			sx *= 10
		}
		for sy <= y {
			sy *= 10
		}
		return x*sy+y > y*sx+x
	})
	if nums[0] == 0 {
		return "0"
	}
	ans := ""
	for _, n := range nums {
		ans += strconv.Itoa(n)
	}
	return ans
}

func main() {
	nums := []int{3, 30, 34, 5, 9}
	ans := largestNumber(nums)
	fmt.Printf(ans)
}
