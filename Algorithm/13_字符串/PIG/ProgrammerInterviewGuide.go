package main

import "fmt"

func reverseWords(words string) string {
	ans := ""
	for i := len(words) - 1; i >= 0; i-- {
		end := i
		for i > 0 && string(words[i]) != " " {
			i--
		}
		start := i
		ans += words[start:end+1] + " "
	}
	return ans
}

func main() {
	words := "I am a coder"
	ans := reverseWords(words)
	fmt.Println(ans)
}
