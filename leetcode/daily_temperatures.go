package leetcode

/*
739. Daily Temperatures (Medium)
Given an array of integers temperatures represents the daily temperatures,
return an array answer such that answer[i] is the number of days
you have to wait after the ith day to get a warmer temperature.
If there is no future day for which this is possible, keep answer[i] == 0 instead.


Example 1:
Input: temperatures = [73,74,75,71,69,72,76,73]
Output: [1,1,4,2,1,1,0,0]

Example 2:
Input: temperatures = [30,40,50,60]
Output: [1,1,1,0]

Example 3:
Input: temperatures = [30,60,90]
Output: [1,1,0]
*/

// TODO: Try to do solve this problem again later

func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0, len(temperatures))
	output := make([]int, len(temperatures))

	for i := len(temperatures) - 1; i > -1; i-- {
		for len(stack) > 0 && temperatures[i] >= temperatures[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			output[i] = stack[len(stack)-1] - i
		}

		stack = append(stack, i)
	}

	return output
}
