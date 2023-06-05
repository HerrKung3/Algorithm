package main

//https://app.codility.com/demo/results/trainingYFGVMA-7TU/

import (
	"fmt"
	"sort"
)

func Solution(A []int) int {
	sort.Ints(A)
	n := len(A)
	if A[n-1] <= 0 {
		return 1
	} else if A[0] > 1 {
		return 1
	}
	for i := 0; i < n-1; i++ {
		if A[i] <= 0 && A[i+1] > 1 {
			return 1
		}
		if A[i] > 0 && A[i+1] > A[i]+1 {
			return A[i] + 1
		}
	}
	return A[n-1] + 1
}

func main() {
	A := []int{2}
	ans := Solution(A)
	fmt.Println(ans)
}
