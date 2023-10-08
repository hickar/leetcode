package structures

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedListBasic(t *testing.T) {
	list := NewLinkedList[int]()

	list.Add(1)
	assert.Equal(t, 1, list.NextNode.Value)

	list.Add(2)
	assert.NotNil(t, list.NextNode.NextNode)
	assert.Equal(t, 2, list.NextNode.NextNode.Value)
}

func TestLinkedListAsSlice(t *testing.T) {
	tests := []struct {
		input          []int
		expectedOutput []int
	}{
		{
			input:          []int{0},
			expectedOutput: []int{0},
		},
		{
			input:          []int{1, 2, 3},
			expectedOutput: []int{1, 2, 3},
		},
		{
			input:          []int{1, 1, 1, 1},
			expectedOutput: []int{1, 1, 1, 1},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestLinkedListAsSlice_Case%d", i), func(t *testing.T) {
			list := NewLinkedList[int]()

			for _, item := range tt.input {
				list.Add(item)
			}

			actualOutput := list.Slice()
			assert.Equal(t, tt.expectedOutput, actualOutput)
		})
	}
}

func TestLinkedListRemove(t *testing.T) {
	tests := []struct {
		list           *ListNode[int]
		itemToRemove   int
		expectedOutput []int
	}{
		{
			list:           NewLinkedList[int](1, 2, 3),
			itemToRemove:   3,
			expectedOutput: []int{1, 2},
		},
		{
			list:           NewLinkedList[int](1, 2, 3),
			itemToRemove:   2,
			expectedOutput: []int{1, 3},
		},
		{
			list:           NewLinkedList[int](1, 2, 3),
			itemToRemove:   1,
			expectedOutput: []int{2, 3},
		},
		{
			list:           NewLinkedList[int](),
			itemToRemove:   1,
			expectedOutput: []int{},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestLinkedListRemove_Case%d", i), func(t *testing.T) {
			list := tt.list
			list.Remove(tt.itemToRemove)

			actualOutput := list.Slice()
			assert.Equal(t, tt.expectedOutput, actualOutput)
		})
	}
}

func TestLinkedListContains(t *testing.T) {
	tests := []struct {
		list          *ListNode[int]
		item          int
		shouldContain bool
	}{
		{
			list:          NewLinkedList[int](1, 2, 3),
			item:          2,
			shouldContain: true,
		},
		{
			list:          NewLinkedList[int](1, 2, 3),
			item:          3,
			shouldContain: true,
		},
		{
			list:          NewLinkedList[int](1, 2, 3),
			item:          4,
			shouldContain: false,
		},
		{
			list:          NewLinkedList[int](),
			item:          0,
			shouldContain: false,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestLinkedListContains_Case%d", i), func(t *testing.T) {
			contains := tt.list.Contains(tt.item)
			assert.Equal(t, tt.shouldContain, contains)
		})
	}
}

//func TestLinkedListReverse(t *testing.T) {
//	tests := []struct {
//		list     *ListNode[int]
//		expected []int
//	}{
//		{
//			list:     NewLinkedList(1, 2, 3, 4),
//			expected: []int{4, 3, 2, 1},
//		},
//		{
//			list:     NewLinkedList(1),
//			expected: []int{1},
//		},
//	}
//
//	for i, tt := range tests {
//		t.Run(fmt.Sprintf("TestLinkedListReverse_Case%d", i+1), func(t *testing.T) {
//			var actual []int
//			require.NotPanics(t, func() {
//				actual = tt.list.Reverse().Slice()
//			})
//
//			assert.Equal(t, tt.expected, actual)
//		})
//	}
//}
