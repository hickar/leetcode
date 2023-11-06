package search

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	cases := []struct {
		arr            []int
		item           int
		expectedOutput int
	}{
		{
			arr:            []int{1, 2, 3, 4, 5, 6},
			item:           3,
			expectedOutput: 2,
		},
		{
			arr:            []int{1, 2, 3, 4, 5},
			item:           5,
			expectedOutput: 4,
		},
		{
			arr:            []int{1, 2, 3},
			item:           4,
			expectedOutput: -1,
		},
		{
			arr:            []int{1},
			item:           1,
			expectedOutput: 0,
		},
		{
			arr:            []int{0},
			item:           1,
			expectedOutput: -1,
		},
		{
			arr:            []int{},
			item:           100,
			expectedOutput: -1,
		},
	}

	for i, tt := range cases {
		t.Run(fmt.Sprintf("Case%d", i), func(t *testing.T) {
			assert.NotPanics(t, func() {
				actualOutput := BinarySearch(tt.arr, tt.item)
				assert.Equal(t, tt.expectedOutput, actualOutput)
			})
		})
	}
}
