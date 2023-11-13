package leetcode

import "sort"

/*
274. H-Index
Given an array of integers citations where citations[i]
is the number of citations a researcher received for their ith paper,
return the researcher's h-index.
According to the definition of h-index on Wikipedia:
The h-index is defined as the maximum value of h such that the given researcher
has published at least h papers that have each been cited at least h times.

Example 1:
Input: citations = [3,0,6,1,5]
Output: 3
Explanation: [3,0,6,1,5] means the researcher has 5 papers in total
and each of them had received 3, 0, 6, 1, 5 citations respectively.
Since the researcher has 3 papers with at least 3 citations each
and the remaining two with no more than 3 citations each, their h-index is 3.

Example 2:
Input: citations = [1,3,1]
Output: 1

*/

func findHIndex(citations []int) int {
	citationsList := numArray(citations)
	sort.Sort(citationsList)

	hIndex := 0
	for i, paperCitationsCount := range citationsList {
		if hIndex >= len(citations)-i {
			break
		}
		if paperCitationsCount > hIndex {
			hIndex = min(paperCitationsCount, len(citations)-i)
		}
	}

	return hIndex
}

type numArray []int

func (n numArray) Len() int {
	return len(n)
}

func (n numArray) Less(i, j int) bool {
	return n[i] < n[j]
}

func (n numArray) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
