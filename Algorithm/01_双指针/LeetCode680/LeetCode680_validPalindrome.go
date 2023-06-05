package main

import "fmt"

func validPalindrome(s string) bool {
	n := len(s)
	p1, p2 := 0, n-1
	for p1 < p2 {
		if s[p1] == s[p2] {
			p1++
			p2--
		} else {
			return deleteOneChar(s[p1:p2]) || deleteOneChar(s[p1+1:p2+1])
		}
	}
	return true
}

func deleteOneChar(s string) bool {
	fmt.Println(s)
	p1, p2 := 0, len(s)-1
	for p1 < p2 {
		if s[p1] != s[p2] {
			return false
		} else {
			p1++
			p2--
		}
	}
	return true
}

func main() {
	s := "ahbhac"
	ans := validPalindrome(s)
	fmt.Println(ans)
}
