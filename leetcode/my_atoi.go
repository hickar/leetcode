package leetcode

import (
	"math"
	"strings"
)

func MyAtoi(s string) int {
	var (
		result  = 0
		numbers = 0
		current = 0
		sign    = 1
	)

	s = strings.Trim(s, " ")

	if s == "" || s == " " {
		return 0
	}

	if s[0] == '-' {
		sign = -1
		current = 1
	} else if s[0] == '+' {
		sign = 1
		current = 1
	}

	for i := current; i < len(s); i++ {
		r := rune((s)[i])

		if (r >= 65 && r <= 90) ||
			(r >= 97 && r <= 122) {
			break
		}

		if r == '-' || r == '+' || r == ' ' || r == '.' {
			if numbers == 0 {
				return 0
			} else {
				break
			}
		}

		if r >= 48 && r <= 57 {
			result = result*10 + runeToInt(r)
			numbers++
		}

		if result*sign < math.MinInt32 {
			return math.MinInt32
		} else if math.MaxInt32 < result*sign {
			return math.MaxInt32
		}
	}

	return result * sign
}

func runeToInt(r rune) int {
	return int(r - 48)
}
