package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddBinary(t *testing.T) {
	tests := []struct {
		inputA   string
		inputB   string
		expected string
	}{
		{
			inputA:   "11",
			inputB:   "1",
			expected: "100",
		},
		{
			inputA:   "1010",
			inputB:   "1011",
			expected: "10101",
		},
		{
			inputA:   "10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101",
			inputB:   "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011",
			expected: "110111101100010011000101110110100000011101000101011001000011011000001100011110011010010011000000000",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			actual := AddBinary(tt.inputA, tt.inputB)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
