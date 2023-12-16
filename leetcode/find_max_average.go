package leetcode

/*
643. Maximum Average Subarray I (Easy)
You are given an integer array nums consisting of n elements, and an integer k.
Find a contiguous subarray whose length is equal to k
that has the maximum average value and return this value.
Any answer with a calculation error less than 10-5 will be accepted.


Example 1:
Input: nums = [1,12,-5,-6,50,3], k = 4
Output: 12.75000
Explanation: Maximum average is (12 - 5 - 6 + 50) / 4 = 51 / 4 = 12.75

Example 2:
Input: nums = [5], k = 1
Output: 5.00000

*/

func FindMaxAverage(nums []int, k int) float64 {
	sum := calcSum(nums[0:k])
	maxSum := sum

	for low, high := 1, k+1; high < len(nums)+1; {
		sum = sum - nums[low-1] + nums[high-1]
		maxSum = max(maxSum, sum)

		low++
		high++
	}

	return float64(maxSum) / float64(k)
}

func calcSum(nums []int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}

	return sum
}
