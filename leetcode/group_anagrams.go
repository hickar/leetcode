package leetcode

import "sort"

/*
49. Group Anagrams (Medium)
Given an array of strings strs, group the anagrams together. You can return the answer in any order.
An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
typically using all the original letters exactly once.

Example 1:
Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

Example 2:
Input: strs = [""]
Output: [[""]]

Example 3:
Input: strs = ["a"]
Output: [["a"]]

*/

func GroupAnagrams(strs []string) [][]string {
	groups := make(map[string]map[string]int)

	for _, word := range strs {
		tmpSortedWord := runeArray(word)
		sort.Sort(tmpSortedWord)
		sortedWord := string(tmpSortedWord)

		_, ok := groups[sortedWord]
		if !ok {
			groups[sortedWord] = make(map[string]int)
		}

		groups[sortedWord][word] += 1
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		anagrams := make([]string, 0, len(group))

		for anagram, count := range group {
			for i := 0; i < count; i++ {
				anagrams = append(anagrams, anagram)
			}
		}

		result = append(result, anagrams)
	}

	return result
}

type runeArray []rune

func (r runeArray) Len() int {
	return len(r)
}

func (r runeArray) Less(i, j int) bool {
	return r[i] < r[j]
}

func (r runeArray) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
