package leetcode

/*
345. Reverse Vowels of a String (Easy)
Given a string s, reverse only all the vowels in the string and return it.
The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases, more than once.


Example 1:
Input: s = "hello"
Output: "holle"

Example 2:
Input: s = "leetcode"
Output: "leotcede"
*/

func reverseVowels(s string) string {
	sb := make([]byte, len(s))
	vows := make([]byte, 0, len(s))
	var c byte
	ok := false

	for i := 0; i < len(s); i++ {
		c = s[i]
		if _, ok = vowelMap[toLower(c)]; ok {
			vows = append(vows, c)
			continue
		}

		sb[i] = c
	}

	var vowIdx int
	for i := len(sb) - 1; i >= 0; i-- {
		if sb[i] == '\000' {
			sb[i] = vows[vowIdx]
			vowIdx++
		}
	}

	return string(sb)
}

func reverseVowelsOptimized(s string) string {
	var (
		sb                 = []byte(s)
		l                  = 0
		lc                 byte
		r                  = len(s) - 1
		rc                 byte
		isLVowel, isRVowel bool
	)

	for l < r {
		lc, rc = sb[l], sb[r]
		isLVowel = isVowel(lc)
		if !isLVowel {
			l++
		}

		isRVowel = isVowel(rc)
		if !isRVowel {
			r--
		}

		if isLVowel && isRVowel {
			sb[l], sb[r] = sb[r], sb[l]
			l++
			r--
		}
	}

	return string(sb)
}

func isVowel(c byte) bool {
	_, ok := vowelMap[toLower(c)]
	return ok
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + ('a' - 'A')
	}

	return c
}

var vowelMap = map[byte]struct{}{
	'a': {},
	'e': {},
	'o': {},
	'i': {},
	'u': {},
}
