package leetcode

import "strings"

/*
6. Zigzag Conversion (Medium)

The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this:
(you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R

And then read line by line: "PAHNAPLSIIGYIR"
Write the code that will take a string and make this conversion given a number of rows:
string convert(string s, int numRows);


Example 1:
Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"

Example 2:
Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:
P     I    N
A   L S  I G
Y A   H R
P     I

Example 3:
Input: s = "A", numRows = 1
Output: "A"
*/

const rowOffsetBaseline = 2

func convert(s string, numRows int) string {
	if len(s) == 1 || numRows == 1 {
		return s
	}

	var (
		oddRowOffset  = rowOffsetBaseline + rowOffsetBaseline*(numRows-rowOffsetBaseline)
		evenRowOffset = 0
		offset        int
		sb            strings.Builder
	)

	for i := 0; i < numRows; i++ {
		offset = i
		for j := 0; ; j++ {
			switch {
			case j == 0:
				break
			case i == 0 || i == numRows-1:
				offset += max(evenRowOffset, oddRowOffset)
			case j%2 == 0:
				offset += evenRowOffset
			case j%2 == 1:
				offset += oddRowOffset
			}

			if offset >= len(s) {
				break
			}

			sb.WriteByte(s[offset])
		}

		oddRowOffset -= rowOffsetBaseline
		evenRowOffset += rowOffsetBaseline
	}

	return sb.String()
}

/* 2 rows (+2 step, 1 gap)
A C
B D

/* 3 rows (+4 step, 2 gap)
A   F
B D
C


4 (+6 step, 3 gap)
A   G
B  F
C E
D


5 (+8 step, 4 gap)
A      I    Q
B     HJ   PR
C   G  K  O S
D F    L N  T
E      M

6 (+10 step, 5 gap)
A      K
B     J
C    I
D   H
E G
F

7 (+12 step, 6 gap)
A      M
B     L
C    K
D   J
E  I
F H
G
*/
