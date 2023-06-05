package main

import "fmt"

func matrixReshape(mat [][]int, r int, c int) [][]int {
	ans := make([][]int, r)
	for i, _ := range ans {
		ans[i] = make([]int, c)
	}
	M := len(mat)
	N := len(mat[0])

	if M*N != r*c {
		return mat
	}
	for i := 0; i < M*N; i++ {
		row := i / c
		column := i % c
		ans[row][column] = mat[i/N][i%N]
	}
	return ans
}

func main() {
	mat := [][]int{{1, 2}, {3, 4}, {5, 6}}
	ans := matrixReshape(mat, 2, 3)
	fmt.Println(ans)
}
