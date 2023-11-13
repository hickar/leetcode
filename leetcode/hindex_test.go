package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHIndex(t *testing.T) {
	cases := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{3, 0, 6, 1, 5},
			expected: 3,
		},
		{
			input:    []int{1, 3, 1},
			expected: 1,
		},
		{
			input:    []int{100},
			expected: 1,
		},
	}

	for i, tt := range cases {
		t.Run(fmt.Sprintf("Case%d", i+1), func(t *testing.T) {
			actual := findHIndex(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
