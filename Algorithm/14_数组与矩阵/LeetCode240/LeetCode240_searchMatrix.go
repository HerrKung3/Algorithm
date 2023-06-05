package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	M := len(matrix)
	N := len(matrix[0])
	for i, j := N-1, 0; i >= 0 && j < M; {
		num := matrix[j][i]
		if num == target {
			return true
		} else if num < target {
			j++
		} else {
			i--
		}
	}
	return false
}

func main() {
	//matrix := [][]int{
	//	{1, 4, 7, 11, 15},
	//	{2, 5, 8, 12, 19},
	//	{3, 6, 9, 16, 22},
	//	{10, 13, 14, 17, 24},
	//	{18, 21, 23, 26, 30}}
	matrix := [][]int{{1, 1}}
	target := 0
	ans := searchMatrix(matrix, target)
	fmt.Println(ans)

}
