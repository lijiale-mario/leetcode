package heap

import "container/heap"

// Item represents an item in the priority queue.
// You might want to customize this for specific problems.
// For a generic heap, the items themselves would need to be comparable,
// or you'd pass a comparison function.
// For simplicity, this example uses integers.

// IntHeap implements heap.Interface for a min-heap of integers.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // Min-heap
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push adds an element to the heap.
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop removes and returns the minimum element from the heap.
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// NewMinHeap creates a new min-heap and initializes it.
func NewMinHeap() *IntHeap {
	h := &IntHeap{}
	heap.Init(h)
	return h
}

// PushToMinHeap adds an element to the min-heap.
func PushToMinHeap(h *IntHeap, val int) {
	heap.Push(h, val)
}

// PopFromMinHeap removes and returns the minimum element from the min-heap.
// Returns 0 and false if the heap is empty.
func PopFromMinHeap(h *IntHeap) (int, bool) {
	if h.Len() == 0 {
		return 0, false
	}
	return heap.Pop(h).(int), true
}

// PeekMinHeap returns the minimum element without removing it.
// Returns 0 and false if the heap is empty.
func PeekMinHeap(h *IntHeap) (int, bool) {
	if h.Len() == 0 {
		return 0, false
	}
	return (*h)[0], true
}

// MaxHeap implements heap.Interface for a max-heap of integers.
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // Max-heap
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push adds an element to the max-heap.
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop removes and returns the maximum element from the max-heap.
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// NewMaxHeap creates a new max-heap and initializes it.
func NewMaxHeap() *MaxHeap {
	h := &MaxHeap{}
	heap.Init(h)
	return h
}

// PushToMaxHeap adds an element to the max-heap.
func PushToMaxHeap(h *MaxHeap, val int) {
	heap.Push(h, val)
}

// PopFromMaxHeap removes and returns the maximum element from the max-heap.
// Returns 0 and false if the heap is empty.
func PopFromMaxHeap(h *MaxHeap) (int, bool) {
	if h.Len() == 0 {
		return 0, false
	}
	return heap.Pop(h).(int), true
}

// PeekMaxHeap returns the maximum element without removing it.
// Returns 0 and false if the heap is empty.
func PeekMaxHeap(h *MaxHeap) (int, bool) {
	if h.Len() == 0 {
		return 0, false
	}
	return (*h)[0], true
}

// Example of a custom item for a priority queue if you need more than just ints
/*
type PQItem struct {
    Value    string // The value of the item; arbitrary.
    Priority int    // The priority of the item in the queue.
    Index    int    // The index of the item in the heap.
}

// PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*PQItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
    // We want Pop to give us the highest, not lowest, priority so we use greater than here.
    return pq[i].Priority > pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
    pq[i].Index = i
    pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
    n := len(*pq)
    item := x.(*PQItem)
    item.Index = n
    *pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    old[n-1] = nil  // avoid memory leak
    item.Index = -1 // for safety
    *pq = old[0 : n-1]
    return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *PQItem, value string, priority int) {
    item.Value = value
    item.Priority = priority
    heap.Fix(pq, item.Index)
}
*/