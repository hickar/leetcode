package leetcode

import (
	"github.com/Hickar/siaod/structures"
	"sort"
)

type IntegerSlice []int

func (s IntegerSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s IntegerSlice) Len() int {
	return len(s)
}

func (s IntegerSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func MergeLinkedLists(l1, l2 *structures.ListNode[int]) *structures.ListNode[int] {
	var (
		s1 []int
		s2 []int
	)

	for {
		s1 = append(s1, l1.Value)
		if l1.NextNode == nil {
			break
		}

		l1 = l1.NextNode
	}
	for {
		s2 = append(s2, l2.Value)
		if l2.NextNode == nil {
			break
		}

		l2 = l2.NextNode
	}

	resultValues := IntegerSlice(append(s1, s2...))
	sort.Sort(resultValues)

	resultNode := &structures.ListNode[int]{}
	curNode := resultNode

	for i, value := range resultValues {
		curNode.Value = value

		if i < len(resultValues)-1 {
			curNode.NextNode = &structures.ListNode[int]{}
			curNode = curNode.NextNode
		}
	}

	return resultNode
}
