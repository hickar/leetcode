package leetcode

import "strconv"

func IsPalindrome(x int) bool {
	s := strconv.Itoa(x)
	length := len(s) - 1

	for i := 0; i <= length/2; i++ {
		if s[i] != s[length-i] {
			return false
		}
	}

	return true
}
