package main

import "fmt"

//给定一个二维的矩阵，包含 'X' 和 'O'（字母 O）。
//找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。
func main() {
	//X(0,0) X(0,1) X(0,2) X(0,3)
	//X(1,0) O(1,1) O(1,2) X(1,3)
	//X(2,0) X(2,1) O(2,2) X(2,3)
	//X(3,0) O(3,1) X(3,2) X(3,3)
	var board = [][]byte{
		{88, 88, 88, 88, 88},
		{88, 79, 79, 88, 88},
		{88, 88, 88, 88, 88},
		{88, 79, 79, 88, 88},
		{88, 79, 88, 88, 88},
	}
	solve(board)
	fmt.Println(board)
}

var n, m int

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	n, m = len(board), len(board[0])
	for i := 0; i < n; i++ {
		dfs(board, i, 0)
		dfs(board, i, m-1)
	}
	for i := 1; i < m-1; i++ {
		dfs(board, 0, i)
		dfs(board, n-1, i)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func dfs(board [][]byte, x, y int) {
	if x < 0 || x >= n || y < 0 || y >= m || board[x][y] != 'O' {
		return
	}
	board[x][y] = 'A'
	dfs(board, x+1, y)
	dfs(board, x-1, y)
	dfs(board, x, y+1)
	dfs(board, x, y-1)
}
