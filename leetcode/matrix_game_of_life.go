package leetcode

/*
289. Game of Life (Medium)

According to Wikipedia's article:
"The Game of Life, also known simply as Life, is a cellular automaton
devised by the British mathematician John Horton Conway in 1970."

The board is made up of an m x n grid of cells, where each cell has an initial state:
live (represented by a 1) or dead (represented by a 0).
Each cell interacts with its eight neighbors (horizontal, vertical, diagonal)
using the following four rules (taken from the above Wikipedia article):
- Any live cell with fewer than two live neighbors dies as if caused by under-population.
- Any live cell with two or three live neighbors lives on to the next generation.
- Any live cell with more than three live neighbors dies, as if by over-population.
- Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.

The next state is created by applying the above rules simultaneously to every cell in the current state,
where births and deaths occur simultaneously.
Given the current state of the m x n grid board, return the next state.


Example 1:
Input: board = [[0,1,0],[0,0,1],[1,1,1],[0,0,0]]
Output: [[0,0,0],[1,0,1],[0,1,1],[0,1,0]]

Example 2:
Input: board = [[1,1],[1,0]]
Output: [[1,1],[1,1]]


Constraints:
- m == board.length
- n == board[i].length
- 1 <= m, n <= 25
- board[i][j] is 0 or 1.
*/

type cellState struct {
	Row   int
	Col   int
	Value int
}

func gameOfLife(board [][]int) {
	var (
		updates   []cellState
		curValue  int
		nextValue int
	)

	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[0]); col++ {
			curValue = board[row][col]
			nextValue = calcNextValue(board, row, col)
			if curValue != nextValue {
				updates = append(updates, cellState{
					Row:   row,
					Col:   col,
					Value: nextValue,
				})
			}
		}
	}

	for _, update := range updates {
		board[update.Row][update.Col] = update.Value
	}
}

func calcNextValue(board [][]int, row, col int) int {
	var (
		rowStart = row - 1
		rowEnd   = row + 1
	)
	if rowStart == -1 {
		rowStart = row
	}
	if rowEnd == len(board) {
		rowEnd = row
	}

	var (
		colStart = col - 1
		colEnd   = col + 1
	)
	if colStart == -1 {
		colStart = col
	}
	if colEnd == len(board[0]) {
		colEnd = col
	}

	neighbors := 0
	for i := rowStart; i <= rowEnd; i++ {
		for j := colStart; j <= colEnd; j++ {
			if i == row && j == col {
				continue
			}

			if board[i][j] == 1 {
				neighbors++
			}
		}
	}

	if board[row][col] == 0 && neighbors == 3 {
		return 1
	}
	if board[row][col] == 1 && neighbors == 2 || neighbors == 3 {
		return 1
	}

	return 0
}
