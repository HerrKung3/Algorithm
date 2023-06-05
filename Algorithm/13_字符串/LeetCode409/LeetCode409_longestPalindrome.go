package main

import "fmt"

func longestPalindrome(s string) int {
	letterFreq := make([]int, 58)
	for i, _ := range s {
		letterFreq[s[i]-65]++
	}
	//出现奇数次的字母将其频次减1，但回文中间的字符可以是奇数次
	odd := 0
	for _, freq := range letterFreq {
		if freq%2 != 0 {
			odd++
		}
	}
	d := 0
	if odd > 0 {
		d = 1
	}
	return len(s) - odd + d
}

func main() {
	s := "abccccdd"
	ans := longestPalindrome(s)
	fmt.Println(ans)
}
