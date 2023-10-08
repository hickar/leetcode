package structures

type ListNode[T comparable] struct {
	Value    T
	NextNode *ListNode[T]
}

func NewLinkedList[T comparable](items ...T) *ListNode[T] {
	list := &ListNode[T]{}

	for _, item := range items {
		list.Add(item)
	}

	return list
}

func (l *ListNode[T]) Add(item T) {
	curNode := l

	for {
		if curNode.NextNode == nil {
			curNode.NextNode = NewLinkedList[T]()
			curNode.NextNode.Value = item
			break
		}

		curNode = curNode.NextNode
	}
}

func (l *ListNode[T]) Remove(item T) bool {
	curNode := l

	for {
		if curNode.NextNode == nil {
			break
		}

		if curNode.NextNode.Value == item {
			curNode.NextNode = curNode.NextNode.NextNode
			return true
		}

		curNode = curNode.NextNode
	}

	return false
}

func (l *ListNode[T]) Contains(item T) bool {
	curNode := l

	for {
		if curNode != l && curNode.Value == item {
			return true
		}

		lastNode := curNode.NextNode
		if lastNode == nil {
			break
		}

		curNode = lastNode
	}

	return false
}

func (l *ListNode[T]) Slice() []T {
	s := make([]T, 0)
	curNode := l

	for {
		if curNode.NextNode == nil {
			break
		}

		curNode = curNode.NextNode
		s = append(s, curNode.Value)
	}

	return s
}

func (l *ListNode[T]) Len() int {
	var (
		curNode = l
		i       int
	)

	for {
		if curNode.NextNode == nil {
			break
		}

		curNode = curNode.NextNode
		i += 1
	}

	return i
}

func (l *ListNode[T]) Clone() *ListNode[T] {
	return NewLinkedList(l.Slice()...)
}

func (l *ListNode[T]) Reverse() *ListNode[T] {
	var (
		prevNode *ListNode[T]
		curNode  = l
		nextNode = l.NextNode
	)

	for curNode != nil {
		nextNode = curNode.NextNode
		curNode.NextNode = prevNode
		prevNode = curNode

		curNode = nextNode
	}

	return prevNode
}
