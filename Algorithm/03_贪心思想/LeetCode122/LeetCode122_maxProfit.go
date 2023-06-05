package main

import "fmt"

func maxProfit(prices []int) int {
	mp := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1]-prices[i] > 0 {
			mp += prices[i+1] - prices[i]
		}
	}
	return mp
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	ans := maxProfit(prices)
	fmt.Println(ans)
}
