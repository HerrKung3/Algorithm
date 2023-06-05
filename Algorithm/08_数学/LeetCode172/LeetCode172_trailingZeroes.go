package main

import (
	"fmt"
)

func trailingZeroes(n int) int {
	ans := 0
	for i := 5; i <= n; i += 5 {
		for j := i; j%5 == 0; j /= 5 {
			ans++
		}
	}
	return ans
}

func main() {
	ans := trailingZeroes(30)
	fmt.Println(ans)
}
