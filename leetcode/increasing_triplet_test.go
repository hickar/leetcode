package leetcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIncreasingTriplet(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
	}{
		{
			input:    []int{1, 2, 3, 4, 5},
			expected: true,
		},
		{
			input:    []int{5, 4, 3, 2, 1},
			expected: false,
		},
		{
			input:    []int{2, 1, 5, 0, 4, 6},
			expected: true,
		},
		{
			input:    []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			expected: false,
		},
		{
			input:    []int{2, 1, 5, 0, 6},
			expected: true,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Case_%d", i), func(t *testing.T) {
			actual := IncreasingTriplet(tt.input)
			assert.Equal(t, tt.expected, actual, "%v", tt.input)
		})
	}
}
