package leetcode

/*
36. Valid Sudoku (Medium)

Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:
– Each row must contain the digits 1-9 without repetition.
– Each column must contain the digits 1-9 without repetition.
– Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.

Note:
A Sudoku board (partially filled) could be valid but is not necessarily solvable.
Only the filled cells need to be validated according to the mentioned rules.
*/

func IsValidSudoku(board [][]byte) bool {
	rows := make(map[int]map[byte]struct{}, 9)
	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]struct{}, 9)
	}

	cols := make(map[int]map[byte]struct{}, 9)
	for i := 0; i < 9; i++ {
		cols[i] = make(map[byte]struct{}, 9)
	}

	boxes := make(map[int]map[byte]struct{}, 9)
	for i := 0; i < 9; i++ {
		boxes[i] = make(map[byte]struct{}, 9)
	}

	var ok bool
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			cur := board[i][j]
			if cur == '.' {
				continue
			}

			boxNum := getBoxNum(i, j)

			if _, ok = rows[i][cur]; ok {
				return false
			}
			if _, ok = cols[j][cur]; ok {
				return false
			}
			if _, ok = boxes[boxNum][cur]; ok {
				return false
			}

			rows[i][cur] = struct{}{}
			cols[j][cur] = struct{}{}
			boxes[boxNum][cur] = struct{}{}
		}
	}

	return true
}

func getBoxNum(i, j int) int {
	boxX := 1
	switch {
	case i > 2 && i < 6:
		boxX = 2
	case i > 5:
		boxX = 3
	}

	boxY := 0
	switch {
	case j > 2 && j < 6:
		boxY = 1
	case j > 5:
		boxY = 2
	}

	return (boxX + (3 * boxY)) - 1
}
