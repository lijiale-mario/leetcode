package bloomfilter

import (
	"hash/fnv"
	// You might use other hash functions like MurmurHash3 for better distribution
)

// BloomFilter represents the Bloom Filter itself.
type BloomFilter struct {
	bitSet    []bool // The bit array
	m         int    // Size of the bit array
	k         int    // Number of hash functions
	hashFuncs []func(data []byte) uint64 // Slice of hash functions
}

// NewBloomFilter creates a new Bloom Filter.
// m: size of the bit array
// k: number of hash functions
// Note: Choosing optimal m and k depends on the expected number of items and desired false positive rate.
func NewBloomFilter(m, k int) *BloomFilter {
	bf := &BloomFilter{
		bitSet:    make([]bool, m),
		m:         m,
		k:         k,
		hashFuncs: make([]func(data []byte) uint64, k),
	}

	// Initialize k hash functions. For simplicity, we'll use variations of FNV.
	// In a real application, you'd want k independent hash functions.
	// This is a simplified way to generate k different hash outputs from one base hash.
	for i := 0; i < k; i++ {
		seed := uint64(i) // Different seed for each hash function variation
		bf.hashFuncs[i] = func(data []byte) uint64 {
			h := fnv.New64a() // FNV-1a is a good non-cryptographic hash function
			// Incorporate the seed to make hashes different
			// A simple way is to append or prepend seed bytes, or XOR with seed.
			// For this example, let's just use the seed in a basic way.
			// This is NOT a cryptographically secure or perfectly distributed set of k hashes.
			// It's a placeholder for a more robust k-hash generation strategy.
			// A common technique is to use two hash functions h1 and h2, and generate k hashes as h1(x) + i*h2(x).
			// For now, let's just pretend we have k distinct good hashes by varying a seed to FNV.
			// This is a simplification for the skeleton.
			// Consider using something like: `h.Write([]byte{byte(seed)})` before `h.Write(data)`
			// or `return baseHash + seed * anotherHash` if you have two base hashes.
			
			// Simple seed incorporation (example only, might not be ideal for distribution)
			seededData := append(data, byte(seed))
			h.Write(seededData)
			return h.Sum64()
		}
	}
	return bf
}

// Add adds an item (represented as a byte slice) to the Bloom Filter.
func (bf *BloomFilter) Add(item []byte) {
	for i := 0; i < bf.k; i++ {
		hashValue := bf.hashFuncs[i](item)
		index := hashValue % uint64(bf.m)
		bf.bitSet[index] = true
	}
}

// Contains checks if an item (represented as a byte slice) might be in the set.
// Returns true if the item is possibly in the set (could be a false positive).
// Returns false if the item is definitely not in the set.
func (bf *BloomFilter) Contains(item []byte) bool {
	for i := 0; i < bf.k; i++ {
		hashValue := bf.hashFuncs[i](item)
		index := hashValue % uint64(bf.m)
		if !bf.bitSet[index] {
			return false // Definitely not in the set
		}
	}
	return true // Possibly in the set (could be a false positive)
}

// Note: The quality of the hash functions is crucial for a Bloom Filter's performance.
// The FNV hash used here is decent, but for k hash functions, one typically uses
// techniques like double hashing (g_i(x) = h1(x) + i*h2(x)) or enhanced double hashing
// to generate k distinct hash values from one or two base hash functions.
// The example above with simple seeding is a placeholder for a more robust approach.