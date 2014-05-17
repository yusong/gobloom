package gobloom

import (
	"hash"
	"hash/fnv"
	"math"
)

type BloomAttributes struct {
	length uint
	k      uint
	hashfn hash.Hash64
	n      uint
}

func (bf *BloomAttributes) EstimateCurrentFPR() float64 {
	k, n, length := float64(bf.k), float64(bf.n), float64(bf.length)
	return math.Pow((1 - math.Exp(-k*n/length)), k)
}

type BloomFilter struct {
	BloomAttributes
	bitset *BitSet
}

func NewBloomFilter(length uint, k uint) *BloomFilter {
	bf := new(BloomFilter)
	bf.length = length
	bf.k = k
	bf.bitset = NewBitSet(length)
	bf.hashfn = fnv.New64()
	bf.n = 0
	return bf
}

func (bf *BloomFilter) hash(b []byte) (uint, uint) {
	bf.hashfn.Reset()
	bf.hashfn.Write(b)
	hash64 := bf.hashfn.Sum64()
	high32 := uint32(hash64 >> 32)
	low32 := uint32(hash64 & (1<<32 - 1))
	return uint(high32), uint(low32)
}

func (bf *BloomFilter) Add(b []byte) {
	high32, low32 := bf.hash(b)
	for i := uint(0); i < bf.k; i++ {
		index := (high32 + i*low32) % bf.length
		bf.bitset.Set(index)
	}
	bf.n++
}

func (bf *BloomFilter) Contains(b []byte) bool {
	ret := true
	high32, low32 := bf.hash(b)
	for i := uint(0); i < bf.k; i++ {
		index := (high32 + i*low32) % bf.length
		ret = ret && bf.bitset.Has(index)
	}
	return ret
}

func EstimateFPR(ulength uint, uk uint, un uint) float64 {
	length, k, n := float64(ulength), float64(uk), float64(un)
	return math.Pow((1 - math.Exp(-k*n/length)), k)
}

type CountingBloomFilter struct {
	BloomAttributes
	buckets *BitSet
}

func NewCountingBloomFilter(length uint, k uint) *CountingBloomFilter {
	cbf := new(CountingBloomFilter)
	cbf.length = length
	cbf.k = k
	cbf.buckets = NewBitSet(length)
	cbf.hashfn = fnv.New64()
	cbf.n = 0
	return cbf
}

func (bf *CountingBloomFilter) hash(b []byte) (uint, uint) {
	bf.hashfn.Reset()
	bf.hashfn.Write(b)
	hash64 := bf.hashfn.Sum64()
	high32 := uint32(hash64 >> 32)
	low32 := uint32(hash64 & (1<<32 - 1))
	return uint(high32), uint(low32)
}

func (bf *CountingBloomFilter) Add(b []byte) {
	high32, low32 := bf.hash(b)
	for i := uint(0); i < bf.k; i++ {
		index := (high32 + i*low32) % bf.length
		bf.buckets.Add(index)
	}
	bf.n++
}

func (bf *CountingBloomFilter) Contains(b []byte) bool {
	ret := true
	high32, low32 := bf.hash(b)
	for i := uint(0); i < bf.k; i++ {
		index := (high32 + i*low32) % bf.length
		ret = ret && bf.buckets.Has(index)
	}
	return ret
}

func (bf *CountingBloomFilter) Remove(b []byte) {
	high32, low32 := bf.hash(b)
	for i := uint(0); i < bf.k; i++ {
		index := (high32 + i*low32) % bf.length
		bf.buckets.Sub(index)
	}
	bf.n--
}
