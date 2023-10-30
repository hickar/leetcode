package leetcode

/*
205. Isomorphic Strings

Given two strings s and t, determine if they are isomorphic.

Two strings s and t are isomorphic if the characters in s can be replaced to get t.

All occurrences of a character must be replaced with another character while preserving the order of characters.
No two characters may map to the same character, but a character may map to itself.
*/

func IsIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sLetters := make(map[byte]byte)
	tLetters := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		sc, tc := s[i], t[i]

		scc, sok := sLetters[sc]
		tcc, tok := tLetters[tc]
		if (sok && t[i] != scc) || (tok && s[i] != tcc) {
			return false
		}

		sLetters[sc] = tc
		tLetters[tc] = sc
	}

	return true
}
