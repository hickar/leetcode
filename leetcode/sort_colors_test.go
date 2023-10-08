package leetcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortColors(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{2, 0, 2, 1, 1, 0},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			input:    []int{2, 0, 1},
			expected: []int{0, 1, 2},
		},
		{
			input:    []int{1},
			expected: []int{1},
		},
		{
			input:    []int{2, 1},
			expected: []int{1, 2},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			SortColors(tt.input)
			assert.Equal(t, tt.expected, tt.input)
		})
	}
}
