package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMergeSortedArray(t *testing.T) {
	tests := []struct {
		nums1        []int
		nums2        []int
		m            int
		n            int
		expectedNums []int
	}{
		{
			nums1:        []int{1, 2, 3, 0, 0, 0},
			nums2:        []int{2, 5, 6},
			m:            3,
			n:            3,
			expectedNums: []int{1, 2, 2, 3, 5, 6},
		},
		{
			nums1:        []int{1},
			nums2:        []int{},
			m:            1,
			n:            0,
			expectedNums: []int{1},
		},
		{
			nums1:        []int{0},
			nums2:        []int{1},
			m:            0,
			n:            1,
			expectedNums: []int{1},
		},
		{
			nums1:        []int{4, 5, 6, 0, 0, 0},
			nums2:        []int{1, 2, 3},
			m:            3,
			n:            3,
			expectedNums: []int{1, 2, 3, 4, 5, 6},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestMergeSortedArray_Case%d", i+1), func(t *testing.T) {
			require.NotPanics(t, func() {
				MergeSortedArray(tt.nums1, tt.m, tt.nums2, tt.n)
			})
			assert.Equal(t, tt.expectedNums, tt.nums1)
		})
	}
}
