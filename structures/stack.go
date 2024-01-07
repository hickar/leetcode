package structures

type Stack[T any] struct {
	items []T
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{
		items: make([]T, 0, 8),
	}
}

func NewStackWithItems[T any](items []T) Stack[T] {
	if items == nil {
		return Stack[T]{items: []T{}}
	}

	return Stack[T]{items: items}
}

func (s *Stack[T]) Push(items ...T) {
	for _, item := range items {
		s.items = append(s.items, item)
	}
}

func (s *Stack[T]) Pop() T {
	if len(s.items) == 0 {
		panic("trying to pop out of zero-sized stack")
	}

	popItem := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return popItem
}

func (s *Stack[T]) Peek() T {
	if len(s.items) == 0 {
		panic("trying to peek out of zero-sized stack")
	}

	item := s.items[len(s.items)-1]
	return item
}

func (s *Stack[T]) Size() int {
	return len(s.items)
}
