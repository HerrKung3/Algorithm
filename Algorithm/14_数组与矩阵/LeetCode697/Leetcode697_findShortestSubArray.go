package main

import (
	"fmt"
	"math"
)

type entry struct {
	freq  int
	first int
	last  int
}

func findShortestSubArray(nums []int) int {
	//num: [freq, firstIndex, lastIndex]
	//numFreq := make(map[int][]int)
	//for key, _ := range numFreq {
	//	numFreq[key] = make([]int, 3)
	//}
	numFreq := make(map[int]entry)
	for i, num := range nums {
		if e, found := numFreq[num]; found {
			e.freq++
			e.last = i
			numFreq[num] = e
		} else {
			numFreq[num] = entry{1, i, i}
		}
	}
	maxFreq := 0
	minLen := math.MaxInt32
	for k, _ := range numFreq {
		if numFreq[k].freq > maxFreq {
			maxFreq = numFreq[k].freq
			minLen = numFreq[k].last - numFreq[k].first + 1
		} else if numFreq[k].freq == maxFreq {
			minLen = min(minLen, numFreq[k].last-numFreq[k].first+1)
		}
	}
	return minLen
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 2, 2, 3, 1}
	ans := findShortestSubArray(nums)
	fmt.Println(ans)
}
