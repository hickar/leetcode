package leetcode

import (
	"fmt"
	"testing"
)

func TestRomansToInteger(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput int
	}{
		{
			input:          "III",
			expectedOutput: 3,
		},
		{
			input:          "LVIII",
			expectedOutput: 58,
		},
		{
			input:          "MCMXCIV",
			expectedOutput: 1994,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("RomanToInteger_%s=%d", tt.input, tt.expectedOutput), func(t *testing.T) {
			actualOutput := RomanToInteger(tt.input)
			if actualOutput != tt.expectedOutput {
				t.Errorf("expected %s to be %d, got %d", tt.input, tt.expectedOutput, actualOutput)
			}
		})
	}
}
