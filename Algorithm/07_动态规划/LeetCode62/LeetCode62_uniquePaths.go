package main

import "fmt"

func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	dp := make([][]int, m)
	for r, _ := range dp {
		dp[r] = make([]int, n)
	}
	dp[0][0] = 1
	for row, column := 1, 1; row < m || column < n; {
		if column < n {
			dp[0][column] = 1
		}
		if row < m {
			dp[row][0] = 1
		}
		row++
		column++
	}
	//fmt.Println(dp)
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]

}

func main() {
	m := 3
	n := 7
	ans := uniquePaths(m, n)
	fmt.Println(ans)
}
