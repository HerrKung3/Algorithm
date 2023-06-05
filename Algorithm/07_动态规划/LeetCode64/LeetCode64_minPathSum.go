package main

import "fmt"

func minPathSum(grid [][]int) int {
	M := len(grid)
	N := len(grid[0])
	if M == 1 && N == 1 {
		return grid[0][0]
	}
	dp := make([][]int, M)
	for r, _ := range dp {
		dp[r] = make([]int, N)
	}

	dp[0][0] = grid[0][0]
	for c := 1; c < N; c++ {
		dp[0][c] = grid[0][c] + dp[0][c-1]
	}
	for r := 1; r < M; r++ {
		dp[r][0] = grid[r][0] + dp[r-1][0]
	}
	//fmt.Println(dp)
	for r := 1; r < M; r++ {
		for c := 1; c < N; c++ {
			dp[r][c] = min(dp[r-1][c], dp[r][c-1]) + grid[r][c]
		}
	}
	return dp[M-1][N-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1}}
	ans := minPathSum(grid)
	fmt.Println(ans)
}
