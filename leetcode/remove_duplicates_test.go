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
			expectedModifiedInput: []int{1, 2},
		},
		{
			input:                 []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expectedUniqueNumbers: 5,
			expectedModifiedInput: []int{0, 1, 2, 3, 4},
		},
		{
			input:                 []int{},
			expectedUniqueNumbers: 0,
			expectedModifiedInput: []int{},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestRemoveDuplicates_Case%d", i), func(t *testing.T) {
			actualOutput := RemoveDuplicates(tt.input)
			assert.Equal(t, tt.expectedUniqueNumbers, actualOutput)

			for j := 0; j < tt.expectedUniqueNumbers; j++ {
				assert.Equal(t, tt.expectedModifiedInput[j], tt.input[j])
			}
		})
	}
}

func TestMoreThanTwoRemoveDuplicates(t *testing.T) {
	tests := []struct {
		input                   []int
		expectedModifiedInput   []int
		expectedSliceIndexValue int
	}{
		{
			input:                   []int{1, 1, 1, 2, 2, 3},
			expectedModifiedInput:   []int{1, 1, 2, 2, 3},
			expectedSliceIndexValue: 5,
		},
		{
			input:                   []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			expectedModifiedInput:   []int{0, 0, 1, 1, 2, 3, 3},
			expectedSliceIndexValue: 7,
		},
		{
			input:                   []int{1, 1, 1},
			expectedModifiedInput:   []int{1, 1},
			expectedSliceIndexValue: 2,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestRemoveDuplicates_Case%d", i), func(t *testing.T) {
			actualOutput := RemoveMoreThanTwoDuplicates(tt.input)
			assert.Equal(t, tt.expectedSliceIndexValue, actualOutput)

			for j := 0; j < tt.expectedSliceIndexValue; j++ {
				assert.Equal(t, tt.expectedModifiedInput[j], tt.input[j])
			}
		})
	}
}
