package main

import "fmt"

func longestConsecutive(nums []int) int {
	hashMap := make(map[int]bool)
	for _, num := range nums {
		hashMap[num] = true
	}
	lc := 0
	count := 1
	for num, _ := range hashMap {
		if val, found := hashMap[num-1]; !found || !val {
			for hashMap[num+1] {
				count++
				num++
			}
			lc = max(lc, count)
			count = 1
		}
	}
	return lc
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	ans := longestConsecutive(nums)
	fmt.Println(ans)
}
