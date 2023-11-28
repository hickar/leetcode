package leetcode

/*
46. Permutations (Medium)

Given an array nums of distinct integers, return all the possible permutations.
You can return the answer in any order.

Example 1:
Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

Example 2:
Input: nums = [0,1]
Output: [[0,1],[1,0]]

Example 3:
Input: nums = [1]
Output: [[1]]
*/

func Permute(nums []int) [][]int {
	combos := make([][]int, 0, factorial(len(nums)))

	var makeCombos func([]int, []int, []bool)

	makeCombos = func(combo []int, numsLeft []int, numsUsed []bool) {
		if len(combo) == len(nums) {
			combos = append(combos, combo)
			return
		}

		for i, num := range numsLeft {
			if numsUsed[i] {
				continue
			}

			newUsed := append(numsUsed[:0:0], numsUsed...)
			newUsed[i] = true
			makeCombos(append(combo, num), numsLeft, newUsed)
		}
	}

	makeCombos([]int{}, nums, make([]bool, len(nums)))

	return combos
}

func factorial(n int) int {
	acc := n

	for i := n - 1; i > 0; i-- {
		acc *= n
	}

	return acc
}
