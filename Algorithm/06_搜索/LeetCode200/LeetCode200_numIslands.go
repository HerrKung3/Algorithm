package main

import (
	"fmt"
)

func numIslands(grid [][]byte) int {
	M := len(grid)
	N := len(grid[0])
	path := make([][]bool, M)
	for r, _ := range path {
		path[r] = make([]bool, N)
	}
	ans := 0
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if grid[i][j] == '1' && !path[i][j] {
				ans += dfs(grid, i, j, path)
			}
		}
	}
	return ans
}

func dfs(grid [][]byte, i, j int, path [][]bool) int {
	if grid[i][j] == '0' {
		return 0
	}
	path[i][j] = true
	next := [][]int{{i + 1, j}, {i, j + 1}, {i - 1, j}, {i, j - 1}}
	for _, n := range next {
		if 0 <= n[0] && n[0] < len(grid) && 0 <= n[1] && n[1] < len(grid[0]) && !path[n[0]][n[1]] {
			dfs(grid, n[0], n[1], path)
		}
	}
	return 1
}

func main() {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '1'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	ans := numIslands(grid)
	fmt.Println(ans)
}
