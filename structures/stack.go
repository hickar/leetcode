package structures

type Stack[T any] struct {
	items []T
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{
		items: make([]T, 0, 5),
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

func (s *Stack[T]) Pop() *T {
	if len(s.items) == 0 {
		return nil
	}

	popItem := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return &popItem
}

func (s *Stack[T]) Len() int {
	return len(s.items)
}
