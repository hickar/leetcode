package leetcode

import "math"

/*
209. Minimum Size Subarray Sum (Medium)
Given an array of positive integers nums and a positive integer target,
return the minimal length of a subarray whose sum is greater than or equal to target.
If there is no such subarray, return 0 instead.


Example 1:
Input: target = 7, nums = [2,3,1,2,4,3]
Output: 2
Explanation: The subarray [4,3] has the minimal length under the problem constraint.

Example 2:
Input: target = 4, nums = [1,4,4]
Output: 1

Example 3:
Input: target = 11, nums = [1,1,1,1,1,1,1,1]
Output: 0


Constraints:
– 1 <= target <= 109
– 1 <= nums.length <= 105
– 1 <= nums[i] <= 104
*/

func minSubArrayLen(target int, nums []int) int {
	var (
		l      int
		r      int
		curSum = nums[l]
		minLen = math.MaxInt64
	)

	for l <= r && r < len(nums) {
		if curSum >= target {
			minLen = min(minLen, r-l+1)
			if minLen == 1 {
				break
			}
		}

		if r < len(nums)-1 && ((curSum-nums[l]+nums[r+1] <= target) || (l == r)) {
			r++
			curSum += nums[r]
		} else {
			curSum -= nums[l]
			l++
		}
	}

	if minLen == math.MaxInt64 {
		return 0
	}

	return minLen
}
