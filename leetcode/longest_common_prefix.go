package leetcode

import (
	"sort"
	"strings"
)

type Words []string

func (w Words) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

func (w Words) Less(i, j int) bool {
	return w[i] < w[j]
}

func (w Words) Len() int {
	return len(w)
}

func Min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func LongestCommonPrefix(strs []string) string {
	words := Words(strs)
	sort.Sort(words)

	first := words[0]
	last := words[len(words)-1]

	var sb strings.Builder

	for i := 0; i < Min(len(first), len(last)); i++ {
		if first[i] != last[i] {
			break
		}

		sb.WriteRune(rune(first[i]))
	}

	return sb.String()
}
