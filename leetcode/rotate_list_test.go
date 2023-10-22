package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateList(t *testing.T) {
	cases := []struct {
		inputNums        []int
		inputRotateCount int
		expectedNums     []int
	}{
		{
			inputNums:        []int{1, 2, 3, 4, 5, 6, 7},
			inputRotateCount: 3,
			expectedNums:     []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			inputNums:        []int{-1, -100, 3, 99},
			inputRotateCount: 2,
			expectedNums:     []int{3, 99, -1, -100},
		},
		{
			inputNums:        []int{1, 2},
			inputRotateCount: 3,
			expectedNums:     []int{2, 1},
		},
	}

	for i, tt := range cases {
		t.Run(fmt.Sprintf("Case%d", i), func(t *testing.T) {
			RotateList(tt.inputNums, tt.inputRotateCount)
			assert.Equal(t, tt.expectedNums, tt.inputNums)
		})
	}
}
