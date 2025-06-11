package bit

// BIT (Binary Indexed Tree) or Fenwick Tree.
// Used for efficient prefix sum queries and point updates.
// Assumes 1-based indexing for internal operations for easier formula application,
// but can be wrapped for 0-based indexing for the user.

type BIT struct {
	Tree []int // The BIT array
	Size int   // Size of the array it represents (number of elements)
}

// NewBIT creates a new BIT of a given size (for elements 0 to size-1).
// The internal tree will be size+1 to handle 1-based indexing.
func NewBIT(size int) *BIT {
	return &BIT{
		Tree: make([]int, size+1),
		Size: size,
	}
}

// Update adds 'delta' to the element at 'index' (0-based).
func (b *BIT) Update(index int, delta int) {
	if index < 0 || index >= b.Size {
		// Handle error or panic: index out of bounds
		return
	}
	i := index + 1 // Convert to 1-based index for BIT operations
	for i <= b.Size { // Or i <= len(b.Tree) -1
		b.Tree[i] += delta
		i += i & (-i) // Add the last set bit (i.e., i & -i)
	}
}

// Query returns the prefix sum up to 'index' (0-based), i.e., sum(arr[0]...arr[index]).
func (b *BIT) Query(index int) int {
	if index < 0 {
        return 0 // Or handle error
    }
    if index >= b.Size {
        index = b.Size - 1 // Query sum of all elements if index is too large
    }

	sum := 0
	i := index + 1 // Convert to 1-based index
	for i > 0 {
		sum += b.Tree[i]
		i -= i & (-i) // Subtract the last set bit
	}
	return sum
}

// QueryRange returns the sum of elements in the range [left, right] (0-based inclusive).
func (b *BIT) QueryRange(left, right int) int {
	if left < 0 || right >= b.Size || left > right {
        // Handle error or return appropriate value (e.g., 0)
        return 0
    }
	if left == 0 {
		return b.Query(right)
	}
	return b.Query(right) - b.Query(left-1)
}

// BuildBITFromSlice creates a BIT from an existing slice of numbers.
func BuildBITFromSlice(nums []int) *BIT {
    size := len(nums)
    bit := NewBIT(size)
    for i := 0; i < size; i++ {
        bit.Update(i, nums[i])
    }
    return bit
}