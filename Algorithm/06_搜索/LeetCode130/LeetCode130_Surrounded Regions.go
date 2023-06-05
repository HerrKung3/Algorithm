package main

import "fmt"

var M, N int

func solve(board [][]byte) {
	M = len(board)
	N = len(board[0])
	if M <= 2 || N <= 2 {
		return
	}
	for i := 0; i < N; i++ {
		dfs(board, 0, i)
		dfs(board, M-1, i)
	}
	for j := 1; j < M-1; j++ {
		dfs(board, j, 0)
		dfs(board, j, N-1)
	}
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func dfs(board [][]byte, row, column int) {
	if row < 0 || row >= M || column < 0 || column >= N || board[row][column] != 'O' {
		return
	}
	board[row][column] = 'A'
	dfs(board, row+1, column)
	dfs(board, row-1, column)
	dfs(board, row, column+1)
	dfs(board, row, column-1)
}

func main() {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}

	solve(board)
	fmt.Println(board)
}
