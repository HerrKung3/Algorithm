package main

//https://app.codility.com/demo/results/trainingWXSUGX-2HM/

import "fmt"

func Solution(A int, B int, K int) int {
	ans := (B - A) / K
	if A%K == 0 || B%K == 0 {
		ans++
	}
	if A < K && K < B && (B-(B/K*K)+(K-A)) < K {
		ans++
	}
	return ans
}

func main() {
	ans := Solution(3, 13, 4)
	fmt.Println(ans)
}
