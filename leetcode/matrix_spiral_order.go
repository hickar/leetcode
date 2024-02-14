package leetcode

/*
54. Spiral Matrix (Medium)
Given an m x n matrix, return all elements of the matrix in spiral order.


Example 1:
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,3,6,9,8,7,4,5]

Example 2:
Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]


Constraints:
- m == matrix.length
- n == matrix[i].length
- 1 <= m, n <= 10
- -100 <= matrix[i][j] <= 100
*/

func spiralOrder(matrix [][]int) []int {
	output := make([]int, 0, len(matrix)*len(matrix[0]))
	rowCount := len(matrix)
	colCount := len(matrix[0])

	readLeft := func(x, y, steps int) (int, int) {
		for i := 0; i < steps; i++ {
			output = append(output, matrix[y][x-i])
		}
		return x - steps + 1, y
	}

	readRight := func(x, y, steps int) (int, int) {
		for i := 0; i < steps; i++ {
			output = append(output, matrix[y][x+i])
		}
		return x + steps - 1, y
	}

	readDown := func(x, y, steps int) (int, int) {
		for i := 0; i < steps; i++ {
			output = append(output, matrix[y+i][x])
		}
		return x, y + steps - 1
	}

	readUp := func(x, y, steps int) (int, int) {
		for i := 0; i < steps; i++ {
			output = append(output, matrix[y-i][x])
		}
		return x, y - steps + 1
	}

	// Read values from matrix OUTER layer
	readRight(0, 0, colCount)
	if rowCount == 1 {
		return output
	}
	readDown(colCount-1, 1, rowCount-1)
	if colCount == 1 {
		return output
	}
	readLeft(colCount-2, rowCount-1, colCount-1)

	// Init loop for reading matrix values in spiral order
	var (
		xi, yi     = colCount - 2, rowCount - 2
		curX, curY = 0, len(matrix) - 2
	)

	for {
		if yi <= 0 {
			break
		}

		curX, curY = readUp(curX, curY, yi)
		curX++
		yi--

		if xi <= 0 {
			break
		}
		curX, curY = readRight(curX, curY, xi)
		curY++
		xi--

		if yi <= 0 {
			break
		}
		curX, curY = readDown(curX, curY, yi)
		curX--
		yi--

		if xi <= 0 {
			break
		}
		curX, curY = readLeft(curX, curY, xi)
		curY--
		xi--
	}

	return output
}
