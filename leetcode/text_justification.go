package leetcode

import (
	"math"
	"strings"
)

/*
68. Text Justification (Hard)

Given an array of strings words and a width maxWidth,
format the text such that each line has exactly maxWidth characters and is fully (left and right) justified.
You should pack your words in a greedy approach; that is, pack as many words as you can in each line.
Pad extra spaces ' ' when necessary so that each line has exactly maxWidth characters.
Extra spaces between words should be distributed as evenly as possible.
If the number of spaces on a line does not divide evenly between words,
the empty slots on the left will be assigned more spaces than the slots on the right.
For the last line of text, it should be left-justified, and no extra space is inserted between words.

Note:
- A word is defined as a character sequence consisting of non-space characters only.
- Each word's length is guaranteed to be greater than 0 and not exceed maxWidth.
- The input array words contains at least one word.


Example 1:
Input: words = ["This", "is", "an", "example", "of", "text", "justification."], maxWidth = 16
Output:
[
   "This    is    an",
   "example  of text",
   "justification.  "
]

Example 2:
Input: words = ["What","must","be","acknowledgment","shall","be"], maxWidth = 16
Output:
[
  "What   must   be",
  "acknowledgment  ",
  "shall be        "
]
Explanation: Note that the last line is "shall be    " instead of "shall     be", because the last line must be left-justified instead of fully-justified.
Note that the second line is also left-justified because it contains only one word.

Example 3:
Input: words = ["Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we","do"], maxWidth = 20
Output:
[
  "Science  is  what we",
  "understand      well",
  "enough to explain to",
  "a  computer.  Art is",
  "everything  else  we",
  "do                  "
]


Constraints:
- 1 <= words.length <= 300
- 1 <= words[i].length <= 20
- words[i] consists of only English letters and symbols.
- 1 <= maxWidth <= 100
- words[i].length <= maxWidth
*/

func fullJustify(words []string, maxWidth int) []string {
	var (
		lines            []string
		curLine          []string
		curLineCharCount int
	)

	for i, word := range words {
		spaceCount := len(curLine)

		if curLineCharCount+len(word)+spaceCount > maxWidth {
			lines = append(
				lines,
				justifyLine(curLine, curLineCharCount, maxWidth),
			)
			curLine = []string{}
			curLineCharCount = 0
		}

		curLine = append(curLine, word)
		curLineCharCount += len(word)

		if i == len(words)-1 {
			lines = append(
				lines,
				justifyLeft(curLine, curLineCharCount),
			)
			break
		}
	}

	return lines
}

func justifyLine(words []string, charCount, maxWidth int) string {
	var (
		remainingSpaces = maxWidth - charCount
		remainingWords  = float64(len(words) - 1)
		buf             strings.Builder
	)

	for i, word := range words {
		buf.WriteString(word)
		if i == len(words)-1 {
			break
		}

		curSpaces := int(math.Ceil(float64(remainingSpaces) / remainingWords))
		remainingSpaces -= curSpaces
		remainingWords -= 1.0

		buf.WriteString(strings.Repeat(" ", curSpaces))
	}

	if remainingSpaces > 0 {
		buf.WriteString(strings.Repeat(" ", int(remainingSpaces)))
	}

	return buf.String()
}

func justifyLeft(words []string, maxWidth int) string {
	var buf strings.Builder

	for i, word := range words {
		buf.WriteString(word)
		if i == len(words)-1 {
			break
		}

		buf.WriteString(" ")
	}

	remainingSpaces := maxWidth - buf.Len()
	if remainingSpaces > 0 {
		buf.WriteString(strings.Repeat(" ", remainingSpaces))
	}

	return buf.String()
}
