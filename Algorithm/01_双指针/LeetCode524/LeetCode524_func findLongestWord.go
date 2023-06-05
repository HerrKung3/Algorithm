package main

import "fmt"

func findLongestWord(s string, dictionary []string) string {
	ans := ""
	for i, _ := range dictionary {
		letter := dictionary[i]
		for j, k := 0, 0; j < len(s) && k < len(letter); {
			if k == len(letter)-1 && letter[k] == s[j] {
				ans = max(ans, letter)
				break
			} else if s[j] == letter[k] {
				j++
				k++
			} else {
				j++
			}
		}
	}
	return ans
}

func max(a, b string) string {
	if len(a) > len(b) {
		return a
	} else if len(a) == len(b) {
		if a < b {
			return a
		} else {
			return b
		}
	}
	return b
}

func main() {
	s := "abpcplea"
	dict := []string{"ale", "apple", "monkey", "plea"}
	ans := findLongestWord(s, dict)
	fmt.Println(ans)
}
