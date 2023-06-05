package main

//https://app.codility.com/demo/results/trainingK3YKHK-2ZC/

import "fmt"

func Solution(A []int) int {
	xor := 0
	for i := 0; i < len(A); i++ {
		xor ^= (i + 1) ^ A[i]
	}
	if xor == 0 {
		return 1
	}
	return 0
}

func main() {
	A := []int{3, 1, 2, 4}
	ans := Solution(A)
	fmt.Println(ans)
}
