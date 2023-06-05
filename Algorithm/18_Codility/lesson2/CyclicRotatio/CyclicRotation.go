package main

//https://app.codility.com/demo/results/trainingDHXUQC-DFB/

import "fmt"

func Solution(A []int, K int) []int {
	n := len(A)
	if n == 0 || K%n == 0 {
		return A
	}
	ans := make([]int, 2*n-K%n)
	for i := 0; i < len(ans); i++ {
		ans[i] = A[i%n]
	}
	return ans[len(ans)-n:]
}

func main() {
	nums := []int{9, 3, 7, 6, 4}
	ans := Solution(nums, 3)
	fmt.Println(ans)
}
