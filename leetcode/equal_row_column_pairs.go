package leetcode

import (
	"strconv"
	"strings"
)

/*
2352. Equal Row and Column Pairs (Medium)

Given a 0-indexed n x n integer matrix grid, return the number of pairs (ri, cj)
such that row ri and column cj are equal.
A row and column pair is considered equal if they contain
the same elements in the same order (i.e., an equal array).


Example 1:
Input: grid = [[3,2,1],[1,7,6],[2,7,7]]
Output: 1
Explanation: There is 1 equal row and column pair:
- (Row 2, Column 1): [2,7,7]

Example 2:
Input: grid = [[3,1,2,2],[1,4,4,5],[2,4,2,2],[2,4,2,2]]
Output: 3
Explanation: There are 3 equal row and column pairs:
- (Row 0, Column 0): [3,1,2,2]
- (Row 2, Column 2): [2,4,2,2]
- (Row 3, Column 2): [2,4,2,2]

*/

func EqualRowColumnPairs(grid [][]int) int {
	rowKeys := make(map[string]int, len(grid))
	colKeys := make(map[string]int, len(grid))

	var key string
	var keyCount int
	for i := 0; i < len(grid); i++ {
		key = rowKey(grid, i)
		keyCount, _ = rowKeys[key]
		rowKeys[key] = keyCount + 1

		key = colKey(grid, i)
		keyCount, _ = colKeys[key]
		colKeys[key] = keyCount + 1
	}

	var pairsCount int
	var keyPairs int
	var otherKeyCount int
	var ok bool

	for rKey, rCount := range rowKeys {
		otherKeyCount, ok = colKeys[rKey]
		if !ok {
			continue
		}

		keyPairs = rCount * otherKeyCount

		rowKeys[rKey] = rCount - keyPairs
		colKeys[rKey] = otherKeyCount - keyPairs
		pairsCount += keyPairs
	}

	return pairsCount
}

func colKey(grid [][]int, col int) string {
	acc := make([]int, len(grid))

	for i := 0; i < len(grid); i++ {
		acc = append(acc, grid[i][col])
	}

	return intToStr(acc)
}

func rowKey(grid [][]int, row int) string {
	acc := make([]int, len(grid))

	for i := 0; i < len(grid); i++ {
		acc = append(acc, grid[row][i])
	}

	return intToStr(acc)
}

func intToStr(nums []int) string {
	strs := make([]string, len(nums))

	for i, num := range nums {
		strs[i] = strconv.Itoa(num)
	}

	return strings.Join(strs, " ")
}
