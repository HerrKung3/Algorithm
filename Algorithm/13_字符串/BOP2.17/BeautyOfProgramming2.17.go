package main

import (
	"fmt"
	"strings"
)

func stringRecycle(s string, k int) string {
	n := len(s)
	s1 := s[:n-k]
	s2 := s[n-k:]
	s1 = reversal(strings.Split(s1, ""))
	s2 = reversal(strings.Split(s2, ""))
	return reversal(strings.Split(s1+s2, ""))
}

func reversal(s []string) string {
	n := len(s)
	for i := 0; i < n/2; i++ {
		s[i], s[n-i-1] = s[n-i-1], s[i]
	}
	return strings.Join(s, "")
}

func main() {
	ans := stringRecycle("abcd123", 3)
	fmt.Println(ans)
}
