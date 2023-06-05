package main

//https://app.codility.com/demo/results/trainingUUDV7U-5K2/
//77%

import (
	"fmt"
)

func Solution(N int, A []int) []int {
	counters := make([]int, N)
	for _, n := range A {
		if 1 <= n && n <= N {
			counters[n-1]++
		} else if n == N+1 {
			maxCounter(counters)
		}
	}
	return counters
}

func maxCounter(counters []int) {
	max := 0
	for _, n := range counters {
		if n > max {
			max = n
		}
	}
	for i, _ := range counters {
		counters[i] = max
	}
}

func main() {
	A := []int{3, 4, 4, 6, 1, 4, 4}
	ans := Solution(5, A)
	fmt.Println(ans)
}
