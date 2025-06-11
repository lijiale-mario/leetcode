package queue

// Queue represents a generic queue.
// For LeetCode, it's often implemented using a slice.
type Queue[T any] struct {
	items []T
}

// NewQueue creates a new queue.
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{items: []T{}}
}

// Enqueue adds an item to the end of the queue.
func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the item from the front of the queue.
// Returns the zero value of T and false if the queue is empty.
func (q *Queue[T]) Dequeue() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

// Peek returns the item from the front of the queue without removing it.
// Returns the zero value of T and false if the queue is empty.
func (q *Queue[T]) Peek() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}
	return q.items[0], true
}

// IsEmpty checks if the queue is empty.
func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

// Size returns the number of items in the queue.
func (q *Queue[T]) Size() int {
	return len(q.items)
}