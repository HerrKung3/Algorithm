package main

//https://app.codility.com/demo/results/trainingP7GVF5-F5A/

import (
	"fmt"
	"math"
)

func Solution(A []int) int {
	sum := 0
	for _, num := range A {
		sum += num
	}
	min := math.MaxInt32
	leftSum, rightSum := 0, sum
	for i := 0; i < len(A)-1; i++ {
		leftSum += A[i]
		rightSum = sum - leftSum
		diff := abs(rightSum, leftSum)
		if diff < min {
			min = diff
		}
	}
	return min
}

func abs(a, b int) int {
	if a-b < 0 {
		return b - a
	}
	return a - b
}

func main() {
	//A := []int{3, 1, 2, 4, 3}
	A := []int{-1000, 1000}
	ans := Solution(A)
	fmt.Println(ans)
}
