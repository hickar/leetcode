package leetcode

/*
3. Longest Substring Without Repeating Characters (Medium)
Given a string s, find the length of the longest substring
without repeating characters.


Example 1:
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

Example 2:
Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Example 3:
Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	var (
		letters         = make(map[byte]int)
		l, r, k, maxLen int
		c               byte
		ok              bool
	)

	for r < len(s) {
		c = s[r]

		if k, ok = letters[c]; ok {
			delete(letters, c)

			for l <= k {
				l++
			}
		}

		letters[c] = r

		maxLen = max(
			maxLen,
			r-l+1,
		)
		r++
	}

	return maxLen
}
