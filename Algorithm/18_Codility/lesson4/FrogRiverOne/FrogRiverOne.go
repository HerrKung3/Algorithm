package main

//https://app.codility.com/demo/results/trainingWTU6JJ-JQK/
//63%
import (
	"fmt"
	"math"
)

func Solution(X int, A []int) int {
	length := len(A)
	if length == 1 && X == 1 {
		return 0
	}
	path := make(map[int]bool)
	mask := 0
	for i := 0; i < length; i++ {
		if value, found := path[A[i]]; !found || !value {
			path[A[i]] = true
			offset := 1 << (X - A[i])
			mask += offset
		}
		if i >= X && (mask&(int(math.Pow(2, float64(X)))-1) == (int(math.Pow(2, float64(X))) - 1)) {
			return i
		}
	}
	return -1
}

func main() {
	A := []int{1, 3, 1, 4, 2, 3, 4, 5}
	ans := Solution(5, A)
	fmt.Println(ans)
}
