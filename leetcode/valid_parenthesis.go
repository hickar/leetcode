package leetcode

import (
	"github.com/Hickar/siaod/structures"
)

func AreValidParenthesis(str string) bool {
	if len(str)%2 != 0 {
		return false
	}

	stack := structures.NewStack[rune]()

	for _, c := range str {
		switch c {
		case '{', '(', '[':
			stack.Push(c)
			continue
		case '}':
			poppedValue := stack.Pop()
			if poppedValue == nil || *poppedValue != '{' {
				return false
			}
			continue
		case ']':
			poppedValue := stack.Pop()
			if poppedValue == nil || *poppedValue != '[' {
				return false
			}
			continue
		case ')':
			poppedValue := stack.Pop()
			if poppedValue == nil || *poppedValue != '(' {
				return false
			}
			continue
		}
	}

	return stack.Len() == 0
}
