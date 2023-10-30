package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		input    string
		expected bool
	}{
		{
			input:    "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			input:    "race a car",
			expected: false,
		},
		{
			input:    " ",
			expected: true,
		},
	}

	for i, tt := range cases {
		t.Run(fmt.Sprintf("Case%d", i), func(t *testing.T) {
			actual := IsPalindrome(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
