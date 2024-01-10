package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack_Basic(t *testing.T) {
	stack := NewStack[int]()

	stack.Push(1, 2, 3)

	stackLen := stack.Size()
	assert.Equal(t, 3, stackLen)

	poppedItems := make([]int, 0, stack.Size())
	for stack.Size() > 0 {
		poppedItems = append(poppedItems, stack.Pop())
	}

	assert.Equal(t, []int{3, 2, 1}, poppedItems)
}

func TestStack_PopZeroLength(t *testing.T) {
	stack := NewStack[int]()

	assert.Panics(t, func() {
		stack.Pop()
	})
}
