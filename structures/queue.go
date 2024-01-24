package structures

type Queue[T comparable] struct {
	items []T
}

func NewQueue[T comparable](cap int) *Queue[T] {
	return &Queue[T]{
		items: make([]T, 0, cap),
	}
}

func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.items) == 0 {
		return *new(T), false
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

func (q *Queue[T]) Len() int {
	return len(q.items)
}

func (q *Queue[T]) RemoveAt(i int) (T, bool) {
	if i < 0 || i > len(q.items)-1 {
		return *new(T), false
	}

	item := q.items[i]
	q.items = append(q.items[:i], q.items[i+1:]...)
	return item, true
}

func (q *Queue[T]) IndexOf(key T) int {
	for i, item := range q.items {
		if item == key {
			return i
		}
	}

	return -1
}
