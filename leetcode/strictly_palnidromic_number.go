package leetcode

/*
2396. Strictly Palindromic Number (Medium)
An integer n is strictly palindromic if, for every base b between 2 and n - 2 (inclusive),
the string representation of the integer n in base b is palindromic.
Given an integer n, return true if n is strictly palindromic and false otherwise.
A string is palindromic if it reads the same forward and backward.


Example 1:
Input: n = 9
Output: false
Explanation: In base 2: 9 = 1001 (base 2), which is palindromic.
In base 3: 9 = 100 (base 3), which is not palindromic.
Therefore, 9 is not strictly palindromic so we return false.
Note that in bases 4, 5, 6, and 7, n = 9 is also not palindromic.

Example 2:
Input: n = 4
Output: false
Explanation: We only consider base 2: 4 = 100 (base 2), which is not palindromic.
Therefore, we return false.
*/

// LOL
func isStrictlyPalindromicV2(n int) bool {
	return false
}

func isStrictlyPalindromic(n int) bool {
	for i := 2; i < 10; i++ {
		if !isPalindrome(toUnitsBase(n, i)) {
			return false
		}
	}

	return true
}

func toUnitsBase(n, base int) []int {
	var bits []int

	for n > 0 {
		bits = append(bits, n%base)
		n /= base
	}

	return reverse(bits)
}

func isPalindrome(n []int) bool {
	for i := 0; i < len(n)/2; i++ {
		if n[i] != n[len(n)-1-i] {
			return false
		}
	}

	return true
}

func reverse(s []int) []int {
	ss := make([]int, len(s))

	for i := 0; i < len(s); i++ {
		ss[i] = s[len(s)-1-i]
	}

	return ss
}
