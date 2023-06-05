package main

import (
	"fmt"
	"strings"
)

// 判断能否通过s1移位得到s2
func isContained(s1, s2 string) bool {
	return strings.Contains(s1+s1, s2)
}

func main() {
	s1 := "Hello"
	s2 := "lloH"
	ans := isContained(s1, s2)
	fmt.Println(ans)
}
