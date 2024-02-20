package leetcode

import "sort"

/*
40. Combination Sum II (Medium)
Given a collection of candidate numbers (candidates) and a target number (target),
find all unique combinations in candidates where the candidate numbers sum to target.
Each number in candidates may only be used once in the combination.
Note: The solution set must not contain duplicate combinations.


Example 1:
Input: candidates = [10,1,2,7,6,1,5], target = 8
Output:
[
	[1,1,6],
	[1,2,5],
	[1,7],
	[2,6]
]

Example 2:
Input: candidates = [2,5,2,1,2], target = 5
Output:
[
	[1,2,2],
	[5]
]


Constraints:
- 1 <= candidates.length <= 100
- 1 <= candidates[i] <= 50
- 1 <= target <= 30
*/

func combinationSum2(candidates []int, target int) [][]int {
	var combos [][]int

	var findCombos func([]int, []int, int, int)

	sort.Ints(candidates)

	findCombos = func(cands, cur []int, step, sum int) {
		if sum >= target || len(cands) == 0 {
			if sum == target {
				tmp := make([]int, len(cur))
				copy(tmp, cur)
				combos = append(combos, tmp)
			}

			return
		}

		for i := step; i < len(cands); i++ {
			if i > step && cands[i] == cands[i-1] {
				continue
			}

			findCombos(
				cands,
				append(cur, cands[i]),
				i+1,
				sum+cands[i],
			)
		}
	}

	findCombos(candidates, nil, 0, 0)

	return combos
}
