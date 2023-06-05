package main

import (
	"fmt"
	"strings"
)

func reverseVowels(s string) string {
	t := []byte(s)
	n := len(t)
	i, j := 0, n-1
	for i < j {
		for i < j && !strings.Contains("aeiouAEIOU", string(t[i])) {
			i++
		}
		for j > i && !strings.Contains("aeiouAEIOU", string(t[j])) {
			j--
		}
		t[i], t[j] = t[j], t[i]
		i++
		j--

	}
	return string(t)
}

func main() {
	ans := reverseVowels("lEetcode")
	fmt.Println(ans)
}
