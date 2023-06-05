package main

import "fmt"

var dirs = [][]int{{0, 1}, {-1, 0}, {0, -1}, {1, 0}}
var m, n int

func pacificAtlantic(heights [][]int) [][]int {
	m, n = len(heights), len(heights[0])
	pacific := make([][]bool, m)
	atlantic := make([][]bool, m)
	ans := make([][]int, 0)
	for row := range pacific {
		pacific[row] = make([]bool, n)
		atlantic[row] = make([]bool, n)
	}

	for row := 0; row < m; row++ {
		dfs(row, 0, pacific, heights)
	}
	for column := 1; column < n; column++ {
		dfs(0, column, pacific, heights)
	}
	for row := 0; row < m; row++ {
		dfs(row, n-1, atlantic, heights)
	}
	for column := 0; column < n-1; column++ {
		dfs(m-1, column, atlantic, heights)
	}

	for i, row := range pacific {
		for j, val := range row {
			if val && atlantic[i][j] {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return ans
}

func dfs(i, j int, ocean [][]bool, heights [][]int) {
	if ocean[i][j] {
		return
	}
	//只要是能从pacific边界搜索进去的点，就将pacific中改点设置为true，表示可以边界遍历到此处，也表示雨水可以从此处流到pacific中
	//Atlantic同理
	ocean[i][j] = true
	for index := range dirs {
		nx := dirs[index][0] + i
		ny := dirs[index][1] + j
		if 0 <= nx && nx < m && 0 <= ny && ny < n && heights[i][j] <= heights[nx][ny] {
			dfs(nx, ny, ocean, heights)
		}
	}
}

func main() {
	heights := [][]int{
		{1, 2, 2, 3, 5},
		{3, 2, 3, 4, 4},
		{2, 4, 5, 3, 1},
		{6, 7, 1, 4, 5},
		{5, 1, 1, 2, 4}}
	ans := pacificAtlantic(heights)
	fmt.Println(ans)
}
