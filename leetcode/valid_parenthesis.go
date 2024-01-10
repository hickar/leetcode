package leetcode

import (
	"github.com/hickar/leetcode/structures"
)

func AreValidParenthesis(str string) bool {
	if len(str)%2 != 0 {
		return false
	}

	var (
		stack       = structures.NewStack[rune]()
		ok          bool
		poppedValue rune
	)

	for _, c := range str {
		switch c {
		case '{', '(', '[':
			stack.Push(c)
			continue
		case '}':
			poppedValue, ok = stack.Pop()
			if !ok || poppedValue != '{' {
				return false
			}
			continue
		case ']':
			poppedValue, ok = stack.Pop()
			if !ok || poppedValue != '[' {
				return false
			}
			continue
		case ')':
			poppedValue, ok = stack.Pop()
			if !ok || poppedValue != '(' {
				return false
			}
			continue
		}
	}

	return stack.Size() == 0
}
