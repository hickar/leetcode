package structures

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashSetInsertMany(t *testing.T) {
	cases := []struct {
		input       []IntHashable
		expectedLen int
		expectedCap int
	}{
		{
			input:       ToIntHashables(1, 2, 3, 4, 5, 6, 7, 8),
			expectedLen: 8,
			expectedCap: 16,
		},
		{
			input:       ToIntHashables(1, 1, 1, 1, 1, 1, 1),
			expectedLen: 1,
			expectedCap: 4,
		},
	}

	for i, tt := range cases {
		t.Run(fmt.Sprintf("Case%d", i), func(t *testing.T) {
			set := NewHashSet[IntHashable]()
			set.InsertMany(tt.input...)
			assert.Equal(t, tt.expectedLen, set.Len())
			assert.Equal(t, tt.expectedCap, set.Cap())
		})
	}
}

func TestHashSetInsert(t *testing.T) {
	set := NewHashSet[IntHashable]()
	assert.True(t, set.Insert(1))
	assert.False(t, set.Insert(1))
	assert.True(t, set.Insert(2))
}

func TestHashSetRemove(t *testing.T) {
	set := NewHashSet[IntHashable]()
	set.InsertMany(1, 2, 3, 4, 5)

	assert.True(t, set.Remove(5))
	assert.Equal(t, 4, set.Len())

	assert.False(t, set.Remove(5))
	assert.Equal(t, 4, set.Len())
}

func TestHashSet_Contains(t *testing.T) {
	cases := []struct {
		input    []IntHashable
		key      IntHashable
		expected bool
	}{
		{
			input:    ToIntHashables(1, 2, 3, 4),
			key:      3,
			expected: true,
		},
		{
			input:    ToIntHashables(1, 2, 3, 4),
			key:      5,
			expected: false,
		},
		{
			input:    ToIntHashables(),
			key:      0,
			expected: false,
		},
	}

	for i, tt := range cases {
		t.Run(fmt.Sprintf("Case%d", i), func(t *testing.T) {
			set := NewHashSet[IntHashable]()
			set.InsertMany(tt.input...)
			actual := set.Remove(tt.key)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func ToIntHashables(nums ...int) []IntHashable {
	items := make([]IntHashable, 0, len(nums))

	for _, num := range nums {
		items = append(items, IntHashable(num))
	}

	return items
}
