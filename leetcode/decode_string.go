package leetcode

import (
	"strconv"
	"strings"
)

/*
394. Decode String (Medium)

Given an encoded string, return its decoded string.
The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets
is being repeated exactly k times.
Note that k is guaranteed to be a positive integer.
You may assume that the input string is always valid;
there are no extra white spaces, square brackets are well-formed, etc.
Furthermore, you may assume that the original data does not contain any digits
and that digits are only for those repeat numbers, k.
For example, there will not be input like 3a or 2[4].
The test cases are generated so that the length of the output will never exceed 105.


Example 1:
Input: s = "3[a]2[bc]"
Output: "aaabcbc"

Example 2:
Input: s = "3[a2[c]]"
Output: "accaccacc"

Example 3:
Input: s = "2[abc]3[cd]ef"
Output: "abcabccdcdcdef"
*/

func decodeString(s string) string {
	var sb strings.Builder
	var str string
	var c byte

	for i := 0; i < len(s); {
		c = s[i]
		if isDigit(c) {
			str = parseInt(s[i:])
			i += len(str) - 1

			str, i = parseExpression(s, i+2, strToInt(str))
			sb.WriteString(str)
			continue
		}
		sb.WriteByte(c)
		i++
	}

	return sb.String()
}

func parseExpression(s string, i int, n int) (string, int) {
	var sb strings.Builder
	var str string
	var c byte

	for i < len(s) {
		c = s[i]
		if c == ']' {
			i++
			break
		}
		if isDigit(c) {
			str = parseInt(s[i:])
			i += len(str) - 1

			str, i = parseExpression(s, i+2, strToInt(str))
			sb.WriteString(str)
			continue
		}

		sb.WriteByte(c)
		i++
	}

	return strings.Repeat(sb.String(), n), i
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func parseInt(str string) string {
	var (
		acc = make([]byte, 0, 3)
		c   byte
	)

	for i := 0; i < len(str); i++ {
		c = str[i]
		if isDigit(c) {
			acc = append(acc, c)
			continue
		}

		break
	}

	return string(acc)
}

func strToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic("failed to convert '" + str + "' to int")
	}

	return num
}
