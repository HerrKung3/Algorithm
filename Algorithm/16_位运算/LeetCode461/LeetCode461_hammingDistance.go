package main

import (
	"fmt"
	"math/bits"
)

func hammingDistance(x int, y int) int {
	return bits.OnesCount(uint(x ^ y))
}

func main() {
	ans := hammingDistance(1, 4)
	fmt.Println(ans)
}
