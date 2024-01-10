package leetcode

import "sort"

/*
1657. Determine if Two Strings Are Close (Medium)

Two strings are considered close if you can attain one from the other using the following operations:
    Operation 1: Swap any two existing characters.
        For example, abcde -> aecdb
    Operation 2: Transform every occurrence of one existing character into another existing character,
	and do the same with the other character.
        For example, aacabb -> bbcbaa (all a's turn into b's, and all b's turn into a's)

You can use the operations on either string as many times as necessary.
Given two strings, word1 and word2, return true if word1 and word2 are close, and false otherwise.



Example 1:
Input: word1 = "abc", word2 = "bca"
Output: true
Explanation: You can attain word2 from word1 in 2 operations.
Apply Operation 1: "abc" -> "acb"
Apply Operation 1: "acb" -> "bca"

Example 2:
Input: word1 = "a", word2 = "aa"
Output: false
Explanation: It is impossible to attain word2 from word1, or vice versa, in any number of operations.

Example 3:
Input: word1 = "cabbba", word2 = "abbccc"
Output: true
Explanation: You can attain word2 from word1 in 3 operations.
Apply Operation 1: "cabbba" -> "caabbb"
Apply Operation 2: "caabbb" -> "baaccc"
Apply Operation 2: "baaccc" -> "abbccc"

*/

func AreCloseStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	word1c := []byte(word1)
	sort.Slice(word1c, func(i, j int) bool { return word1c[i] < word1c[j] })
	word1 = string(word1c)

	word2c := []byte(word2)
	sort.Slice(word2c, func(i, j int) bool { return word2c[i] < word2c[j] })
	word2 = string(word2c)

	// "abc", "acb" => "abc", "abc"
	if word1 == word2 {
		return true
	}

	// "cabbba", "abbccc" => "abbccc", "aabbbc"
	letterCount1 := make(map[byte]int, len(word1))
	letterCount2 := make(map[byte]int, len(word2))
	var oc int
	for i := 0; i < len(word1c); i++ {
		oc, _ = letterCount1[word1[i]]
		letterCount1[word1[i]] = oc + 1

		oc, _ = letterCount2[word2[i]]
		letterCount2[word2[i]] = oc + 1
	}
	if len(letterCount1) != len(letterCount2) {
		return false
	}
	if !haveSameKeys(letterCount1, letterCount2) {
		return false
	}

	values1, values2 := values(letterCount1), values(letterCount2)
	sort.Ints(values1)
	sort.Ints(values2)

	return areEqual(values1, values2)
}

func values(m map[byte]int) []int {
	vals := make([]int, 0, len(m))

	for _, value := range m {
		vals = append(vals, value)
	}

	return vals
}

func haveSameKeys(m, mm map[byte]int) bool {
	var ok bool
	for key := range m {
		if _, ok = mm[key]; !ok {
			return false
		}
	}

	for key := range mm {
		if _, ok = m[key]; !ok {
			return false
		}
	}

	return true
}

func areEqual(s, ss []int) bool {
	if len(s) != len(ss) {
		return false
	}

	for i := 0; i < len(s); i++ {
		if s[i] != ss[i] {
			return false
		}
	}

	return true
}
