package leetcode

/*
242. Valid Anagram

Given two strings s and t, return true if t is an anagram of s, and false otherwise.
An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

Example 1:
Input: s = "anagram", t = "nagaram"
Output: true

Example 2:
Input: s = "rat", t = "car"
Output: false
*/

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	letterMap := make(map[rune]int, len(s))

	for _, c := range s {
		letterMap[c] += 1
	}

	for _, c := range t {
		sc, ok := letterMap[c]
		if !ok || sc == 0 {
			return false
		}

		letterMap[c] -= 1
	}

	return true
}
