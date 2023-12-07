package leetcode

/*
1207. Unique Number of Occurrences (Easy)

Given an array of integers arr, return true if the number of occurrences
of each value in the array is unique or false otherwise.

Example 1:
Input: arr = [1,2,2,1,1,3]
Output: true
Explanation: The value 1 has 3 occurrences, 2 has 2 and 3 has 1. No two values have the same number of occurrences.

Example 2:
Input: arr = [1,2]
Output: false

Example 3:
Input: arr = [-3,0,1,-3,1,1,1,-3,10,0]
Output: true

*/

func UniqueOccurrences(arr []int) bool {
	occurrenceMap := make(map[int]int, len(arr))

	var ok bool
	var occurrences int
	for _, num := range arr {
		occurrences, _ = occurrenceMap[num]
		occurrenceMap[num] = occurrences + 1
	}

	unique := make(map[int]struct{}, len(occurrenceMap))
	for _, occurrences := range occurrenceMap {
		_, ok = unique[occurrences]
		if ok {
			return false
		}

		unique[occurrences] = struct{}{}
	}

	return true
}
