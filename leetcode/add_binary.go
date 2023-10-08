package leetcode

import (
	"strings"
)

func AddBinary(a, b string) string {
	if len(a) < len(b) {
		a = padWithZeroes(a, len(b)-len(a))
	} else {
		b = padWithZeroes(b, len(a)-len(b))
	}

	var (
		result    = make([]rune, len(a)+1)
		remainder int
	)

	for i := len(a) - 1; i >= 0; i-- {
		sum := int(a[i]-48) + int(b[i]-48)
		sum, remainder = normalizeBinary(sum, remainder)

		result[i+1] = rune(sum) + 48
		if i == 0 {
			break
		}
	}
	if remainder == 1 {
		result[0] = '1'
	} else {
		result = result[1:]
	}

	return string(result)
}

func normalizeBinary(num, remainder int) (int, int) {
	switch {
	case num == 2 && remainder == 1:
		return 1, 1
	case num == 2:
		return 0, 1
	case num == 1 && remainder == 1:
		return 0, 1
	default:
		return num + remainder, 0
	}
}

func padWithZeroes(s string, padding int) string {
	return strings.Repeat("0", padding) + s
}
