package structures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_Basic(t *testing.T) {
	stack := NewStack[int]()

	stack.Push(1, 2, 3)

	stackLen := stack.Len()
	assert.Equal(t, 3, stackLen)

	poppedItems := make([]int, 0, stack.Len())
	for stack.Len() > 0 {
		poppedItems = append(poppedItems, *stack.Pop())
	}

	assert.Equal(t, []int{3, 2, 1}, poppedItems)
}

func TestStack_PopZeroLength(t *testing.T) {
	stack := NewStack[int]()

	popped := stack.Pop()
	assert.Nil(t, popped)
}
