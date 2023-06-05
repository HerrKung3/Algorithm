package main

import "fmt"

func convertToTitle(columnNumber int) string {
	var ans []byte
	for columnNumber > 0 {
		digit := columnNumber % 26
		if digit == 0 {
			ans = append(ans, byte('A'+25))
			columnNumber = columnNumber/26 - 1
		} else {
			ans = append(ans, byte('A'+digit-1))
			columnNumber /= 26
		}
	}
	reversal(ans)
	return string(ans)
}

func reversal(s []byte) {
	n := len(s)
	for i := 0; i < n/2; i++ {
		s[i], s[n-i-1] = s[n-i-1], s[i]
	}
}

func main() {
	columnNumber := 28
	ans := convertToTitle(columnNumber)
	fmt.Println(ans)
}
