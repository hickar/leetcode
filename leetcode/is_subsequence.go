package leetcode

/*
392. Is Subsequence

Given two strings s and t, return true if s is a subsequence of t, or false otherwise.
A subsequence of a string is a new string that is formed from the original string
by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters.
(i.e., "ace" is a subsequence of "abcde" while "aec" is not).

Example 1:
Input: s = "abc", t = "ahbgdc"
Output: true

Example 2:
Input: s = "axc", t = "ahbgdc"
Output: false

*/

func IsSubsequence(s string, t string) bool {
	var matchCount int

	if s == "" {
		return true
	}

	for _, c := range t {
		if c == rune(s[matchCount]) {
			matchCount += 1
			if matchCount == len(s) {
				break
			}
		}
	}

	return matchCount == len(s)
}
