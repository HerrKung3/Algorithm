package main

import "fmt"

func countBinarySubstrings(s string) int {
	ptr, n, last := 0, len(s), 0
	ans := 0
	for ptr < n {
		c := s[ptr]
		count := 0
		for ptr < n && s[ptr] == c {
			ptr++
			count++
		}
		ans += min(count, last)
		last = count
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	s := "0001100110"
	ans := countBinarySubstrings(s)
	fmt.Println(ans)
}
