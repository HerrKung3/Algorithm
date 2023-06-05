package main

import "fmt"

func maxAreaOfIsland(grid [][]int) int {
	M := len(grid)
	N := len(grid[0])
	ans := 0
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if grid[i][j] == 1 {
				ans = max(ans, dfs(grid, i, j))
			}
		}
	}
	return ans
}

func dfs(grid [][]int, i, j int) int {
	if grid[i][j] == 0 {
		return 0
	}
	ans := 1
	grid[i][j] = 0
	next := [][]int{{i + 1, j}, {i, j + 1}, {i - 1, j}, {i, j - 1}}
	for _, n := range next {
		if 0 <= n[0] && n[0] < len(grid) && 0 <= n[1] && n[1] < len(grid[0]) {
			ans += dfs(grid, n[0], n[1])
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	grid := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}

	ans := maxAreaOfIsland(grid)
	fmt.Println(ans)
}
