package main

import "fmt"

func mySqrt(x int) int {
	left, right, ans := 0, x, -1
	for left <= right {
		mid := (right-left)/2 + left
		//二分查找中注意边界的处理
		if mid*mid <= x {
			ans = mid
			left = mid + 1
		} else if mid*mid > x {
			right = mid - 1
		}
	}
	return ans
}

func main() {
	ans := mySqrt(8)
	fmt.Println(ans)
}
