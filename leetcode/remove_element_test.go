package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		input                 []int
		element               int
		expectedUniqueNumbers int
	}{
		{
			input:                 []int{3, 2, 2, 3},
			element:               3,
			expectedUniqueNumbers: 2,
		},
		{
			input:                 []int{0, 1, 2, 2, 3, 0, 4, 2},
			element:               2,
			expectedUniqueNumbers: 5,
		},
		{
			input:                 []int{},
			element:               0,
			expectedUniqueNumbers: 0,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestRemoveDuplicates_Case%d", i), func(t *testing.T) {
			actualOutput := RemoveElement(tt.input, tt.element)
			assert.Equal(t, tt.expectedUniqueNumbers, actualOutput)
		})
	}
}
