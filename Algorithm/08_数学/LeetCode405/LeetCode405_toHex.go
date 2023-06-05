package main

import (
	"fmt"
	"math"
)

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	negative := num
	if negative < 0 {
		num = math.MaxUint32 + num + 1
	}
	ans := []byte{}
	for num > 0 {
		digit := num % 16
		if digit > 9 {
			ans = append(ans, 'a'+byte(digit-10))
		} else {
			ans = append(ans, '0'+byte(digit))
		}
		num /= 16
	}
	return reversal(ans)
}

func reversal(num []byte) string {
	n := len(num)
	for i := 0; i < n/2; i++ {
		num[i], num[n-i-1] = num[n-i-1], num[i]
	}
	return string(num)
}

func main() {
	num := 26
	ans := toHex(num)
	fmt.Println(ans)
}
