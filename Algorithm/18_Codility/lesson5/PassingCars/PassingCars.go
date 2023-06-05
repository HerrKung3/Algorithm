package main

//https://app.codility.com/demo/results/training5HE4F2-M5D/

import "fmt"

func Solution(A []int) int {
	n := len(A)
	ans := 0
	countZero := 0
	for i := 0; i < n; i++ {
		if A[i] == 0 {
			countZero++
			ans += n - i - countZero
		}
	}
	if ans > 1000000000 {
		return -1
	}
	return ans
}

func main() {
	A := []int{0, 1, 0, 1, 1, 0, 1, 0}
	ans := Solution(A)
	fmt.Println(ans)
}
