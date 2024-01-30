package leetcode

/*
77. Combinations (Medium)
Given two integers n and k, return all possible combinations of k numbers chosen from the range [1, n].
You may return the answer in any order.


Example 1:
Input: n = 4, k = 2
Output: [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]
Explanation: There are 4 choose 2 = 6 total combinations.
Note that combinations are unordered, i.e., [1,2] and [2,1] are considered to be the same combination.

Example 2:
Input: n = 1, k = 1
Output: [[1]]
Explanation: There is 1 choose 1 = 1 total combination.
*/

func combinationsOptimized(n int, k int) [][]int {
	if n == 1 {
		return [][]int{{n}}
	}

	var (
		combos           [][]int
		makeCombinations func([]int, int)
	)

	makeCombinations = func(cur []int, step int) {
		if len(cur) == k {
			combo := make([]int, len(cur))
			copy(combo, cur)
			combos = append(combos, combo)
			return
		}

		for i := step; i <= n; i++ {
			cur = append(cur, i)
			makeCombinations(
				cur,
				i+1,
			)
			cur = cur[:len(cur)-1]
		}
	}

	makeCombinations(nil, 1)

	return combos
}

func combinations(n int, k int) [][]int {
	if n == 1 {
		return [][]int{{n}}
	}

	var combos [][]int

	base := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		base = append(base, i)
	}

	var makeCombinations func([]int, []int, []bool)

	makeCombinations = func(cur, left []int, used []bool) {
		if len(cur) == k {
			combos = append(combos, cur)
			return
		}

		for i, num := range left {
			if used[i] {
				continue
			}

			used[i] = true
			cur = append(cur, num)
			makeCombinations(
				append(append(cur[:0:0], cur...), num),
				left,
				append(used[:0:0], used...),
			)
		}
	}

	makeCombinations(nil, base, make([]bool, n))

	return combos
}
