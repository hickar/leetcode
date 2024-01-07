package leetcode

import "strings"

/*
1768. Merge Strings Alternately (Easy)

You are given two strings word1 and word2. Merge the strings by adding letters in alternating order,
starting with word1. If a string is longer than the other,
append the additional letters onto the end of the merged string.
Return the merged string.


Example 1:
Input: word1 = "abc", word2 = "pqr"
Output: "apbqcr"
Explanation: The merged string will be merged as so:
word1:  a   b   c
word2:    p   q   r
merged: a p b q c r

Example 2:
Input: word1 = "ab", word2 = "pqrs"
Output: "apbqrs"
Explanation: Notice that as word2 is longer, "rs" is appended to the end.
word1:  a   b
word2:    p   q   r   s
merged: a p b q   r   s

Example 3:
Input: word1 = "abcd", word2 = "pq"
Output: "apbqcd"
Explanation: Notice that as word1 is longer, "cd" is appended to the end.
word1:  a   b   c   d
word2:    p   q
merged: a p b q c   d
*/

func mergeAlternately(word1 string, word2 string) string {
	var (
		w1len, w2len = len(word1), len(word2)
		maxlen       = max(w1len, w2len)
		sb           strings.Builder
		i, j         int
	)

	sb.Grow(maxlen)

	for {
		if i > w1len-1 && j > w2len-1 {
			break
		}

		switch {
		case i < j && i < w1len:
			sb.WriteByte(word1[i])
			i++
		case j < i && j < w2len:
			sb.WriteByte(word2[j])
			j++
		case i < w1len:
			sb.WriteByte(word1[i])
			i++
		case j < w2len:
			sb.WriteByte(word2[j])
			j++
		}
	}

	return sb.String()
}
