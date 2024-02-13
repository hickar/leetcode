package leetcode

import "strings"

/*
224. Basic Calculator (Hard)
Given a string s representing a valid expression, implement a basic calculator to evaluate it,
and return the result of the evaluation.
Note: You are not allowed to use any built-in function which evaluates
strings as mathematical expressions, such as eval().


Example 1:
Input: s = "1 + 1"
Output: 2

Example 2:
Input: s = " 2-1 + 2 "
Output: 3

Example 3:
Input: s = "(1+(4+5+2)-3)+(6+8)"
Output: 23


Constraints:
- 1 <= s.length <= 3 * 105
- s consists of digits, '+', '-', '(', ')', and ' '.
- s represents a valid expression.
- '+' is not used as a unary operation (i.e., "+1" and "+(2 + 3)" is invalid).
- '-' could be used as a unary operation (i.e., "-1" and "-(2 + 3)" is valid).
- There will be no two consecutive operators in the input.
- Every number and running calculation will fit in a signed 32-bit integer.
*/

func calculate(s string) int {
	val, _ := expression(strings.TrimSpace(s), 0)
	return val
}

func expression(s string, i int) (int, int) {
	var (
		acc    int
		t      int
		c      byte
		opFunc OpFunc
	)

parseLoop:
	for i < len(s) {
		c = s[i]

		switch {
		case isDigit(c) && opFunc == nil:
			acc, i = term(s, i)

		case c == '+':
			opFunc = OpFunc(opAdd)

		case c == '-':
			opFunc = OpFunc(opSub)

		case isDigit(c) && opFunc != nil:
			t, i = term(s, i)
			acc = opFunc(acc, t)
			opFunc = nil

		case c == '(' && opFunc != nil:
			t, i = expression(s, i+1)
			acc = opFunc(acc, t)
			opFunc = nil

		case c == '(' && opFunc == nil:
			acc, i = expression(s, i+1)

		case c == ')':
			break parseLoop

		case c == ' ':
			break
		}

		i++
	}

	return acc, i
}

func term(s string, i int) (int, int) {
	var (
		operand []byte
		c       byte
	)

parseLoop:
	for i < len(s) {
		switch c = s[i]; {
		case isDigit(c):
			operand = append(operand, c)

		case !isDigit(c) && len(operand) > 0:
			i--
			break parseLoop

		case c == ' ' && len(operand) == 0:
			break
		}

		i++
	}

	return toInt(operand), i
}

type OpFunc func(int, int) int

func opAdd(n1, n2 int) int {
	return n1 + n2
}

func opSub(n1, n2 int) int {
	return n1 - n2
}

func toInt(b []byte) int {
	var acc int

	for i := 0; i < len(b); i++ {
		acc = acc*10 + int(b[i]-'0')
	}

	return acc
}
