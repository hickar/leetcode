package leetcode

/*
2390. Removing Stars From a String (Medium)
You are given a string s, which contains stars *.

In one operation, you can:
    Choose a star in s.
    Remove the closest non-star character to its left, as well as remove the star itself.

Return the string after all stars have been removed.
Note:
    The input will be generated such that the operation is always possible.
    It can be shown that the resulting string will always be unique.


Example 1:
Input: s = "leet**cod*e"
Output: "lecoe"
Explanation: Performing the removals from left to right:
- The closest character to the 1st star is 't' in "leet**cod*e". s becomes "lee*cod*e".
- The closest character to the 2nd star is 'e' in "lee*cod*e". s becomes "lecod*e".
- The closest character to the 3rd star is 'd' in "lecod*e". s becomes "lecoe".
There are no more stars, so we return "lecoe".

Example 2:
Input: s = "erase*****"
Output: ""
Explanation: The entire string is removed, so we return an empty string.
*/

func removeStars(s string) string {
	stack := NewByteStack(len(s))

	var c byte
	for i := 0; i < len(s); i++ {
		c = s[i]
		if c == '*' {
			stack.Pop()
			continue
		}

		stack.Push(c)
	}

	return stack.String()
}

type ByteStack struct {
	items []byte
	size  int
}

func NewByteStack(size int) *ByteStack {
	return &ByteStack{
		items: make([]byte, 0, size),
		size:  0,
	}
}

func (s *ByteStack) Push(c byte) {
	s.items = append(s.items, c)
	s.size++
}

func (s *ByteStack) Pop() byte {
	c := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	s.size--
	return c
}

func (s *ByteStack) String() string {
	strbytes := make([]byte, s.size)

	for i := s.size - 1; i > -1; i-- {
		strbytes[i] = s.Pop()
	}

	return string(strbytes)
}
