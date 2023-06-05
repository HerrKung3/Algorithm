package main

//https://app.codility.com/demo/results/trainingT76YDP-FGH/

import "fmt"

func Solution(A []int) int {
	xor := 0
	for i := 0; i < len(A); i++ {
		xor ^= A[i] ^ (i + 1)
	}
	return xor ^ (len(A) + 1)
}

func main() {
	A := []int{2, 3, 1, 5}
	ans := Solution(A)
	fmt.Println(ans)
}
