package main

import (
	"fmt"
)

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	negative := num < 0
	if negative {
		num = -num
	}
	ans := []byte{}
	for num > 0 {
		ans = append(ans, '0'+byte(num%7))
		num /= 7
	}
	if negative {
		ans = append(ans, '-')
	}
	Reversal(ans)
	return string(ans)
}

func Reversal(num []byte) {
	n := len(num)
	for i := 0; i < len(num)/2; i++ {
		num[i], num[n-i-1] = num[n-i-1], num[i]
	}
}

func main() {
	ans := convertToBase7(100)
	fmt.Println(ans)
}
