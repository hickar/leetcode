package leetcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAreValidParenthesis(t *testing.T) {
	tests := []struct {
		input         string
		shouldBeValid bool
	}{
		{
			input:         "()",
			shouldBeValid: true,
		},
		{
			input:         "()[]{}",
			shouldBeValid: true,
		},
		{
			input:         "(]",
			shouldBeValid: false,
		},
		{
			input:         "([)]",
			shouldBeValid: false,
		},
		{
			input:         "[()]{}",
			shouldBeValid: true,
		},
		{
			input:         "((",
			shouldBeValid: false,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestAreValidParenthesis_Case%d", i), func(t *testing.T) {
			valid := AreValidParenthesis(tt.input)
			assert.Equal(t, tt.shouldBeValid, valid)
		})
	}
}
