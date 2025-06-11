package dsu

// DSU (Disjoint Set Union) or Union-Find data structure.
// This implementation uses path compression and union by size/rank for optimization.

type DSU struct {
	parent []int
	size   []int // or rank
	count  int   // Number of disjoint sets
}

// NewDSU creates a new DSU structure with n elements (0 to n-1).
func NewDSU(n int) *DSU {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return &DSU{
		parent: parent,
		size:   size,
		count:  n,
	}
}

// Find returns the representative (root) of the set containing element x.
// It uses path compression.
func (dsu *DSU) Find(x int) int {
	if dsu.parent[x] == x {
		return x
	}
	dsu.parent[x] = dsu.Find(dsu.parent[x]) // Path compression
	return dsu.parent[x]
}

// Union merges the sets containing elements x and y.
// It uses union by size (or rank).
// Returns true if x and y were in different sets (union happened), false otherwise.
func (dsu *DSU) Union(x, y int) bool {
	rootX := dsu.Find(x)
	rootY := dsu.Find(y)

	if rootX == rootY {
		return false // Already in the same set
	}

	// Union by size: attach smaller tree under root of larger tree
	if dsu.size[rootX] < dsu.size[rootY] {
		dsu.parent[rootX] = rootY
		dsu.size[rootY] += dsu.size[rootX]
	} else {
		dsu.parent[rootY] = rootX
		dsu.size[rootX] += dsu.size[rootY]
	}
	dsu.count--
	return true
}

// Connected checks if elements x and y are in the same set.
func (dsu *DSU) Connected(x, y int) bool {
	return dsu.Find(x) == dsu.Find(y)
}

// Count returns the number of disjoint sets.
func (dsu *DSU) Count() int {
	return dsu.count
}

// SizeOfSet returns the size of the set containing element x.
func (dsu *DSU) SizeOfSet(x int) int {
	rootX := dsu.Find(x)
	return dsu.size[rootX]
}
