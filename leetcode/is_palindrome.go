package leetcode

import (
	"strings"
	"unicode"
)

func IsPalindrome(str string) bool {
	var sb strings.Builder
	sb.Grow(len(str))

	for _, c := range str {
		if !unicode.IsLetter(c) && !unicode.IsDigit(c) {
			continue
		}

		if unicode.IsUpper(c) {
			c = unicode.ToLower(c)
		}

		sb.WriteRune(c)
	}

	s := sb.String()
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}

	return true
}
