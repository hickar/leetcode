package leetcode

/*
59. Spiral Matrix II (Medium)
Given a positive integer n, generate an n x n matrix filled with elements from 1 to n2 in spiral order.


Example 1:
Input: n = 3
Output: [[1,2,3],[8,9,4],[7,6,5]]

Example 2:
Input: n = 1
Output: [[1]]


Constraints:
- 1 <= n <= 20
*/

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	fillUp := func(x, y, steps, n int) (int, int, int) {
		for i := 0; i < steps; i++ {
			matrix[y-i][x] = n
			n++
		}

		return x, y - steps + 1, n
	}

	fillDown := func(x, y, steps, n int) (int, int, int) {
		for i := 0; i < steps; i++ {
			matrix[y+i][x] = n
			n++
		}

		return x, y + steps - 1, n
	}

	fillLeft := func(x, y, steps, n int) (int, int, int) {
		for i := 0; i < steps; i++ {
			matrix[y][x-i] = n
			n++
		}

		return x - steps + 1, y, n
	}

	fillRight := func(x, y, steps, n int) (int, int, int) {
		for i := 0; i < steps; i++ {
			matrix[y][x+i] = n
			n++
		}

		return x + steps - 1, y, n
	}

	cn := 1

	_, _, cn = fillRight(0, 0, n, cn)
	if cn > n*n {
		return matrix
	}
	_, _, cn = fillDown(n-1, 1, n-1, cn)
	if cn > n*n {
		return matrix
	}
	_, _, cn = fillLeft(n-2, n-1, n-1, cn)
	if cn > n*n {
		return matrix
	}

	var (
		curX, curY = 0, n - 2
		xi, yi     = n - 2, n - 2
	)

	for {
		if yi <= 0 {
			break
		}
		curX, curY, cn = fillUp(curX, curY, yi, cn)
		curX++
		yi--

		if xi <= 0 {
			break
		}
		curX, curY, cn = fillRight(curX, curY, xi, cn)
		curY++
		xi--

		if yi <= 0 {
			break
		}
		curX, curY, cn = fillDown(curX, curY, yi, cn)
		curX--
		yi--

		if xi <= 0 {
			break
		}
		curX, curY, cn = fillLeft(curX, curY, xi, cn)
		curY--
		xi--
	}

	return matrix
}
