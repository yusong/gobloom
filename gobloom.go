package gobloom

import (
	"hash"
	"hash/fnv"
	"math"
)

// Base attributes of bloom filter
type BloomAttributes struct {
	length uint        // Size of bloom filter
	k      uint        // Number of hash functions
	hashfn hash.Hash64 // Hash function
	n      uint        // Number of elements in bloom filter
}

// Estimates the false positive rate of current bloom filter.
// According to http://en.wikipedia.org/wiki/Bloom_filter,
// the false positive rate(FPR) of bloom filter with n elements
// is nearly (1-e^(-kn/m))^k, m is the size of bloom filter
func (bf *BloomAttributes) EstimateCurrentFPR() float64 {
	k, n, length := float64(bf.k), float64(bf.n), float64(bf.length)
	return math.Pow((1 - math.Exp(-k*n/length)), k)
}

// Estimates FPR when you have not initialize a bloom fliter object
func EstimateFPR(ulength uint, uk uint, un uint) float64 {
	length, k, n := float64(ulength), float64(uk), float64(un)
	return math.Pow((1 - math.Exp(-k*n/length)), k)
}

// Base bool bloom filter, 1 indicates true and 0 indicates false
type BloomFilter struct {
	BloomAttributes
	bitset *BitSet // Bitset object holds values bits
}

// Returns a new object of Bloom Filter. The length paramter
// indicates size of it, and k means the hash functions it uses.
// It uses fnv1 hash function.
// Reference http://en.wikipedia.org/wiki/Fowler-Noll-Vo_hash_function
func NewBloomFilter(length uint, k uint) *BloomFilter {
	bf := new(BloomFilter)
	bf.length = length
	bf.k = k
	bf.bitset = NewBitSet(length)
	bf.hashfn = fnv.New64()
	bf.n = 0
	return bf
}

// Calls 64-bits fnv1 hash function, the result is uint with
// 64-bits. Cuts off uint64 result and get 2 uint32 values
// the higher 32-bits and lowwer 32-bits.
func (bf *BloomFilter) hash(b []byte) (uint, uint) {
	bf.hashfn.Reset()
	bf.hashfn.Write(b)
	hash64 := bf.hashfn.Sum64()
	high32 := uint32(hash64 >> 32)
	low32 := uint32(hash64 & (1<<32 - 1))
	return uint(high32), uint(low32)
}

// Adds an element into bloom filter. Calls k hash functions
// and sets bits to 1. The k-th hash function is high32*k + low32
func (bf *BloomFilter) Add(b []byte) {
	high32, low32 := bf.hash(b)
	for i := uint(0); i < bf.k; i++ {
		index := (high32*i + low32) % bf.length
		bf.bitset.Set(index)
	}
	bf.n++
}

// Checks whether an element is in bloom filter or not
func (bf *BloomFilter) Contains(b []byte) bool {
	ret := true
	high32, low32 := bf.hash(b)
	for i := uint(0); i < bf.k; i++ {
		index := (high32*i + low32) % bf.length
		ret = ret && bf.bitset.Has(index)
	}
	return ret
}

// Counting bloom filter provides a way to implement a delete
// operation on a Bloom filter. The insert operation is extended
// to increment the value of the buckets and the lookup operation
// checks that each of the required buckets is non-zero. The delete
// operation, obviously, then consists of decrementing the value of
// each of the respective buckets.
type CountingBloomFilter struct {
	BloomAttributes
	buckets *BitSet // Bitset object holds values bits
}

// Returns a new object of Counting Bloom Filter. The length paramter
// indicates size of it, and k means the hash functions it uses.
// It uses fnv1 hash function.
// Reference http://en.wikipedia.org/wiki/Fowler-Noll-Vo_hash_function
func NewCountingBloomFilter(length uint, k uint) *CountingBloomFilter {
	cbf := new(CountingBloomFilter)
	cbf.length = length
	cbf.k = k
	cbf.buckets = NewBitSet(length)
	cbf.hashfn = fnv.New64()
	cbf.n = 0
	return cbf
}

// Calls 64-bits fnv1 hash function, the result is uint with
// 64-bits. Cuts off uint64 result and get 2 uint32 values
// the higher 32-bits and lowwer 32-bits.
func (bf *CountingBloomFilter) hash(b []byte) (uint, uint) {
	bf.hashfn.Reset()
	bf.hashfn.Write(b)
	hash64 := bf.hashfn.Sum64()
	high32 := uint32(hash64 >> 32)
	low32 := uint32(hash64 & (1<<32 - 1))
	return uint(high32), uint(low32)
}

// Adds an element into counting bloom filter. Calls k hash functions
// and increase bits by 1. The k-th hash function is high32*k + low32
func (bf *CountingBloomFilter) Add(b []byte) {
	high32, low32 := bf.hash(b)
	for i := uint(0); i < bf.k; i++ {
		index := (high32*i + low32) % bf.length
		bf.buckets.Add(index)
	}
	bf.n++
}

// Checks whether an element is in counting bloom filter or not
func (bf *CountingBloomFilter) Contains(b []byte) bool {
	ret := true
	high32, low32 := bf.hash(b)
	for i := uint(0); i < bf.k; i++ {
		index := (high32*i + low32) % bf.length
		ret = ret && bf.buckets.Has(index)
	}
	return ret
}

// Removes an element and decrease bits by 1
func (bf *CountingBloomFilter) Remove(b []byte) {
	high32, low32 := bf.hash(b)
	for i := uint(0); i < bf.k; i++ {
		index := (high32*i + low32) % bf.length
		bf.buckets.Sub(index)
	}
	bf.n--
}
