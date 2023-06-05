package main

import (
	"fmt"
)

func Solution(num int) int {
	return num / 2
}

func main() {
	a := 0
	for {
		_, _ = fmt.Scan(&a)
		if a == 0 {
			break
		} else {
			ans := Solution(a)
			fmt.Println(ans)
			//fmt.Printf("%d\n", a)
		}
	}
}
