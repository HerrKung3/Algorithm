package main

import (
	"fmt"
	"strconv"
	"strings"
)

//https://app.codility.com/demo/results/trainingX2RW5C-X8Z/

func Solution(N int) int {
	n := strconv.FormatInt(int64(N), 2)
	last := strings.LastIndex(n, "1")
	gaps := strings.Split(n[:last], "1")
	max := 0
	for _, gap := range gaps {
		if len(gap) > max && strings.Contains(gap, "0") {
			max = len(gap)
		}
	}
	return max
}

func main() {
	ans := Solution(20)
	fmt.Println(ans)
}
