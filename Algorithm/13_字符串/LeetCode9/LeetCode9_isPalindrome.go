package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	} else if x%10 == 0 {
		return false
	}

	num := 0
	for x > num {
		digit := x % 10
		x /= 10
		num = num*10 + digit
		//fmt.Printf("digit=%d, x=%d, num=%d\n", digit, x, num)
	}
	return x == num || x == num/10
}

func main() {
	ans := isPalindrome(12345654321)
	fmt.Println(ans)
}
