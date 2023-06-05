package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	length := len(flowerbed)
	if length == 1 && flowerbed[0] == 0 {
		return n <= 1
	}
	space := 0
	spf := 0
	for i, _ := range flowerbed {
		if (i == 0 && flowerbed[i] == 0) || (i == length-1 && flowerbed[i] == 0) {
			space++
		}
		if flowerbed[i] == 0 {
			space++
		} else {
			curSpf := (space - 1) / 2
			space = 0
			spf += curSpf
		}
	}
	return n <= spf+(space-1)/2
}

func main() {
	flowerbed := []int{1, 0, 0, 0, 0, 1, 0, 1, 0, 0}
	ans := canPlaceFlowers(flowerbed, 2)
	fmt.Println(ans)
}
