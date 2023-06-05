package main

import "fmt"

//数组nums中有n个数，在这n个数的后面添加从0到n的每个整数，则添加了n+1个整数，共有2n+1个整数。
//在2n+1个整数中，丢失的数字只在后面n+1个整数中出现一次，其余的数字在前面n个整数中（即数组中）和后面n+1个整数中各出现一次，即其余的数字都出现了两次

func missingNumber2(nums []int) int {
	xor := 0
	for i, n := range nums {
		xor ^= i ^ n
	}
	return xor ^ len(nums)
}

func main() {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	ans := missingNumber2(nums)
	fmt.Println(ans)
}
