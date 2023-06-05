package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	s1 := make([]int, 26)
	t1 := make([]int, 26)

	for i := 0; i < len(s); i++ {
		s1[s[i]%26]++
		t1[t[i]%26]++
	}
	//fmt.Println(s1)
	//fmt.Println(t1)
	for i := 0; i < len(s1); i++ {
		if s1[i] != t1[i] {
			return false
		}
	}
	return true
}

func main() {
	s := "anagram"
	t := "nagaram"
	ans := isAnagram(s, t)
	fmt.Println(ans)
}
