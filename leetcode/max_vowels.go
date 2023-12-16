package leetcode

/*
1456. Maximum Number of Vowels in a Substring of Given Length (Medium)
Given a string s and an integer k, return the maximum number of vowel letters in any substring of s with length k.
Vowel letters in English are 'a', 'e', 'i', 'o', and 'u'.


Example 1:
Input: s = "abciiidef", k = 3
Output: 3
Explanation: The substring "iii" contains 3 vowel letters.

Example 2:
Input: s = "aeiou", k = 2
Output: 2
Explanation: Any substring of length 2 contains 2 vowels.

Example 3:
Input: s = "leetcode", k = 3
Output: 2
Explanation: "lee", "eet" and "ode" contain 2 vowels.

*/

func MaxVowels(s string, k int) int {
	var ok bool
	var maxCount int

	for i := 0; i < k; i++ {
		if _, ok = vowels[s[i]]; ok {
			maxCount++
		}
	}
	if maxCount == k {
		return maxCount
	}

	curCount := maxCount
	for i, j := 1, k+1; j < len(s)+1; {
		if _, ok = vowels[s[i-1]]; ok {
			curCount--
		}
		if _, ok = vowels[s[j-1]]; ok {
			curCount++
		}

		maxCount = max(maxCount, curCount)
		if maxCount == k {
			return maxCount
		}

		i++
		j++
	}

	return maxCount
}

var vowels = map[byte]struct{}{
	'a': {},
	'e': {},
	'i': {},
	'o': {},
	'u': {},
}
