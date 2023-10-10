package leetcode

import (
	"fmt"
	"testing"

	"github.com/Hickar/siaod/structures"
	"github.com/stretchr/testify/assert"
)

func TestMergeLinkedLists(t *testing.T) {
	tests := []struct {
		lists          [2]*structures.ListNode[int]
		expectedOutput []int
	}{
		{
			lists: [2]*structures.ListNode[int]{
				structures.NewLinkedList(1, 2, 4),
				structures.NewLinkedList(1, 3, 4),
			},
			expectedOutput: []int{1, 1, 2, 3, 4, 4},
		},
		{
			lists: [2]*structures.ListNode[int]{
				structures.NewLinkedList[int](),
				structures.NewLinkedList[int](),
			},
			expectedOutput: []int{},
		},
		{
			lists: [2]*structures.ListNode[int]{
				structures.NewLinkedList[int](),
				structures.NewLinkedList(0),
			},
			expectedOutput: []int{0},
		},
		{
			lists: [2]*structures.ListNode[int]{
				structures.NewLinkedList(4, 5, 6),
				structures.NewLinkedList(1, 2, 3),
			},
			expectedOutput: []int{1, 2, 3, 4, 5, 6},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestMergeLinkedLists_Case%d", i), func(t *testing.T) {
			actualOutput := MergeLinkedLists(tt.lists[0], tt.lists[1]).Slice()
			assert.Equal(t, tt.expectedOutput, actualOutput)
		})
	}
}
