package skiplist

import (
	"math/rand"
	"time"
)

const maxLevel = 16 // Maximum level for the skip list

// Node represents a node in the Skip List.
// For simplicity, this example uses integer values.
// You would typically use a generic type or interface for values.
type Node struct {
	Val     int
	Forward []*Node // Array of forward pointers at different levels
}

// SkipList represents the Skip List itself.
type SkipList struct {
	Header *Node // Header node
	Level  int   // Current maximum level of the skip list (0 to maxLevel-1)
	rand   *rand.Rand
}

// NewNode creates a new skip list node.
func NewNode(val, level int) *Node {
	return &Node{
		Val:     val,
		Forward: make([]*Node, level+1),
	}
}

// NewSkipList creates a new Skip List.
func NewSkipList() *SkipList {
	// Initialize random number generator
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	return &SkipList{
		Header: NewNode(-1, maxLevel-1), // Header value can be sentinel like -1 or MinInt
		Level:  0,
		rand:   r,
	}
}

// randomLevel determines the level for a new node.
// P is the probability of increasing the level (typically 0.25 or 0.5).
func (sl *SkipList) randomLevel() int {
	lvl := 0
	// P = 0.5, so 50% chance to increment level
	for sl.rand.Intn(2) == 0 && lvl < maxLevel-1 {
		lvl++
	}
	return lvl
}

// Search searches for a value in the Skip List.
// Returns true if found, false otherwise.
// (Implementation is non-trivial and omitted for brevity in this skeleton)
func (sl *SkipList) Search(target int) bool {
	// current := sl.Header
	// for i := sl.Level; i >= 0; i-- {
	// 	for current.Forward[i] != nil && current.Forward[i].Val < target {
	// 		current = current.Forward[i]
	// 	}
	// }
	// current = current.Forward[0]
	// return current != nil && current.Val == target
	return false // Placeholder
}

// Insert inserts a value into the Skip List.
// (Implementation is non-trivial and omitted for brevity in this skeleton)
func (sl *SkipList) Insert(val int) {
	// update := make([]*Node, maxLevel)
	// current := sl.Header
	// for i := sl.Level; i >= 0; i-- {
	// 	for current.Forward[i] != nil && current.Forward[i].Val < val {
	// 		current = current.Forward[i]
	// 	}
	// 	update[i] = current
	// }
	// current = current.Forward[0]

	// // If value already exists, you might update it or do nothing
	// if current != nil && current.Val == val {
	// 	return
	// }

	// newLvl := sl.randomLevel()
	// if newLvl > sl.Level {
	// 	for i := sl.Level + 1; i <= newLvl; i++ {
	// 		update[i] = sl.Header
	// 	}
	// 	sl.Level = newLvl
	// }

	// newNode := NewNode(val, newLvl)
	// for i := 0; i <= newLvl; i++ {
	// 	newNode.Forward[i] = update[i].Forward[i]
	// 	update[i].Forward[i] = newNode
	// }
}

// Delete deletes a value from the Skip List.
// (Implementation is non-trivial and omitted for brevity in this skeleton)
func (sl *SkipList) Delete(val int) bool {
	// update := make([]*Node, maxLevel)
	// current := sl.Header
	// for i := sl.Level; i >= 0; i-- {
	// 	for current.Forward[i] != nil && current.Forward[i].Val < val {
	// 		current = current.Forward[i]
	// 	}
	// 	update[i] = current
	// }
	// current = current.Forward[0]

	// if current != nil && current.Val == val {
	// 	for i := 0; i <= sl.Level; i++ {
	// 		if update[i].Forward[i] != current {
	// 			break
	// 		}
	// 		update[i].Forward[i] = current.Forward[i]
	// 	}
	// 	// Remove levels that have no elements
	// 	for sl.Level > 0 && sl.Header.Forward[sl.Level] == nil {
	// 		sl.Level--
	// 	}
	// 	return true
	// }
	return false // Placeholder
}

// Note: A full Skip List implementation requires careful pointer manipulation
// and handling of levels. The above is a structural skeleton.