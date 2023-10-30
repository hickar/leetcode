package leetcode

import "strings"

/*

290. Word Pattern

Given a pattern and a string s, find if s follows the same pattern.
Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in s.

~~~~~ ===== ~~~~~

Example 1:
Input: pattern = "abba", s = "dog cat cat dog"
Output: true

Example 2:
Input: pattern = "abba", s = "dog cat cat fish"
Output: false

*/

func WordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(words) != len(pattern) {
		return false
	}

	letterMap := make(map[string]string)
	wordMap := make(map[string]string)

	for i := 0; i < len(words); i++ {
		lt, wt := string(pattern[i]), words[i]

		ltc, lok := letterMap[lt]
		wtc, wok := wordMap[wt]
		if (lok && wt != ltc) || (wok && lt != wtc) {
			return false
		}

		letterMap[lt] = wt
		wordMap[wt] = lt
	}

	return true
}
