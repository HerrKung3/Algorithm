package main

import (
	"fmt"
	"strconv"
)

func Solution(s string) string {
	ans := 0
	n := len(s)
	for i := 2; i < n; i++ {
		digit := s[i] - 48
		//fmt.Println("digit = ", digit)
		if digit < 10 && i == n-1 {
			ans += int(digit) * helper(10, n-i-1)
		} else if digit < 10 {
			ans += int(digit) * helper(16, n-i-1)
		} else {
			ans += int(digit-7) * helper(16, n-i-1)
		}
	}
	return strconv.Itoa(ans)
}

func helper(n, x int) int {
	if x == 0 {
		return 1
	}
	ans := n
	for i := 0; i < x-1; i++ {
		ans *= n
	}
	return ans
}

func main() {
	a := ""
	_, _ = fmt.Scan(&a)
	ans := Solution(a)
	fmt.Printf("%s\n", ans)

}
