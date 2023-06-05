package main

//https://app.codility.com/demo/results/trainingRFQJDM-6XT/

import "fmt"

func Solution(A []int) int {
	xor := 0
	for _, num := range A {
		xor ^= num
	}
	return xor
}

func main() {
	A := []int{2, 2, 3, 5, 3, 6, 1, 6, 1}
	ans := Solution(A)
	fmt.Println(ans)
}
