package main

import (
	"fmt"
	"strconv"
	"strings"
)

func addBinary(a string, b string) string {
	ans := make([]string, 0)
	flag := 0
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0 || flag == 1; {
		sum := 0
		if i >= 0 {
			sum += int(a[i]) - 48
		}
		if j >= 0 {
			sum += int(b[j]) - 48
		}
		if flag == 1 {
			sum += flag
			flag = 0
		}
		ans = append(ans, strconv.Itoa(sum%2))
		if sum >= 2 {
			flag = 1
		}
		i--
		j--
	}
	return reversal(ans)
}

func reversal(ans []string) string {
	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-i-1] = ans[len(ans)-i-1], ans[i]
	}
	return strings.Join(ans, "")
}

func main() {
	a := "1111"
	b := "1111"
	ans := addBinary(a, b)
	fmt.Println(ans)
}
