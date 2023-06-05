package main

import (
	"fmt"
	"math"
)

func judgeSquareSum(c int) bool {
	j := int(math.Sqrt(float64(c)))
	for i := 0; i <= j; {
		sum := i*i + j*j
		if sum > c {
			j--
		} else if sum < c {
			i++
		} else {
			return true
		}
	}
	return false
}

func main() {
	ans := judgeSquareSum(33)
	fmt.Println(ans)
}
