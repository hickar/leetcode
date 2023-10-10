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

	for curNode.NextNode != nil {
		curNode = curNode.NextNode
	}

	curNode.NextNode = NewLinkedList[T]()
	curNode.NextNode.Value = item
}

func (l *ListNode[T]) Remove(item T) bool {
	curNode := l

	for curNode.NextNode != nil {
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

	for curNode.NextNode != nil {
		curNode = curNode.NextNode
		if curNode.Value == item {
			return true
		}
	}

	return false
}

func (l *ListNode[T]) Slice() []T {
	s := make([]T, 0)
	curNode := l

	for curNode.NextNode != nil {
		curNode = curNode.NextNode
		s = append(s, curNode.Value)
	}

	return s
}

func (l *ListNode[T]) Len() int {
	var i int

	for curNode := l; curNode.NextNode != nil; i++ {
		curNode = curNode.NextNode
	}

	return i
}

func (l *ListNode[T]) Clone() *ListNode[T] {
	return NewLinkedList(l.Slice()...)
}

func (l *ListNode[T]) Reverse() *ListNode[T] {
	var (
		prevNode *ListNode[T]
		curNode  = l.NextNode
		nextNode *ListNode[T]
	)

	for curNode != nil {
		nextNode = curNode.NextNode
		curNode.NextNode = prevNode
		prevNode = curNode

		curNode = nextNode
	}

	return &ListNode[T]{NextNode: prevNode}
}
