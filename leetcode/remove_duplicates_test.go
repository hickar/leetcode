package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		input                 []int
		expectedModifiedInput []int
		expectedUniqueNumbers int
	}{
		{
			input:                 []int{1, 1, 2},
			expectedUniqueNumbers: 2,
		},
		{
			input:                 []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expectedUniqueNumbers: 5,
		},
		{
			input:                 []int{},
			expectedUniqueNumbers: 0,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestRemoveDuplicates_Case%d", i), func(t *testing.T) {
			actualOutput := RemoveDuplicates(tt.input)
			assert.Equal(t, tt.expectedUniqueNumbers, actualOutput)
		})
	}
}
