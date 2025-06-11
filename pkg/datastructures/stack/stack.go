package stack

// Stack represents a generic stack.
// For LeetCode, it's often implemented using a slice.
type Stack[T any] struct {
	items []T
}

// NewStack creates a new stack.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{items: []T{}}
}

// Push adds an item to the top of the stack.
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the item from the top of the stack.
// Returns the zero value of T and false if the stack is empty.
func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}
	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item, true
}

// Peek returns the item from the top of the stack without removing it.
// Returns the zero value of T and false if the stack is empty.
func (s *Stack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

// IsEmpty checks if the stack is empty.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of items in the stack.
func (s *Stack[T]) Size() int {
	return len(s.items)
}