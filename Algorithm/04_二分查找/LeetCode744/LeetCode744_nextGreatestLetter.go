package main

import (
	"fmt"
	"sort"
)

func nextGreatestLetter(letters []byte, target byte) byte {
	n := len(letters)
	if letters[n-1] <= target || letters[0] > target {
		return letters[0]
	}
	return letters[sort.Search(n-1, func(i int) bool {
		return letters[i] > target
	})]
}

func main() {
	letters := []byte{'c', 'f', 'j'}
	ans := nextGreatestLetter(letters, 'd')
	fmt.Println(string(ans))
}
