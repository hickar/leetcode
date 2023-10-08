package leetcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		input          []string
		expectedOutput string
	}{
		{
			input:          []string{"flower", "flow", "flight"},
			expectedOutput: "fl",
		},
		{
			input:          []string{"dog", "racecar", "car"},
			expectedOutput: "",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("LongestCommonPrefix_Case%d", i), func(t *testing.T) {
			actualOutput := LongestCommonPrefix(tt.input)
			assert.Equal(t, tt.expectedOutput, actualOutput)
		})
	}
}
