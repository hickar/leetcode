package leetcode

/*
73. Set Matrix Zeroes (Medium)
Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's.
You must do it in place.


Example 1:
Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
Output: [[1,0,1],[0,0,0],[1,0,1]]

Example 2:
Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]


Constraints:
- m == matrix.length
- n == matrix[0].length
- 1 <= m, n <= 200
- -231 <= matrix[i][j] <= 231 - 1
*/

func setZeroes(matrix [][]int) {
	var (
		rowCount = len(matrix)
		colCount = len(matrix[0])
		rowMap   = make(map[int]struct{}, rowCount)
		colMap   = make(map[int]struct{}, colCount)
	)

	for i := 0; i < rowCount; i++ {
		for j := 0; j < colCount; j++ {
			if matrix[i][j] == 0 {
				rowMap[i] = struct{}{}
				colMap[j] = struct{}{}
			}
		}
	}

	for row := range rowMap {
		for col := 0; col < colCount; col++ {
			matrix[row][col] = 0
		}
	}

	for col := range colMap {
		for row := 0; row < rowCount; row++ {
			matrix[row][col] = 0
		}
	}
}
