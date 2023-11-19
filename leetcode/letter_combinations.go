package leetcode

/*
17. Letter Combinations of a Phone Number (Medium)

Given a string containing digits from 2-9 inclusive,
return all possible letter combinations that the number could represent.
Return the answer in any order.

A mapping of digits to letters (just like on the telephone buttons) is given below.
Note that 1 does not map to any letters.

Example 1:
Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]

Example 2:
Input: digits = ""
Output: []

Example 3:
Input: digits = "2"
Output: ["a","b","c"]

*/

var buttonLetters = map[byte][]byte{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	return calcCombos("", digits)
}

func calcCombos(combo string, nextDigits string) []string {
	if len(nextDigits) == 0 {
		return []string{combo}
	}

	combos := make([]string, 0)
	for _, letter := range buttonLetters[nextDigits[0]] {
		combos = append(combos, calcCombos(combo+string(letter), nextDigits[1:])...)
	}

	return combos
}
