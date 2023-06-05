package main

//https://app.codility.com/demo/results/training7264YW-CET/

import "fmt"

func Solution(X int, Y int, D int) int {
	if (Y-X)/D*D+X == Y {
		return (Y - X) / D
	}
	return (Y-X)/D + 1
}

func main() {
	ans := Solution(85, 85, 30)
	fmt.Println(ans)
}
