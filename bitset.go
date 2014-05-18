package gobloom

// Base data structure of bitset
type BitSet struct {
	length uint   // Size of the bitset
	bits   []uint // Container that holds uint values
}

// Returns a new object of BitSet, the lenght
// parameter indicates the size of BitSet object
func NewBitSet(length uint) *BitSet {
	return &BitSet{length, make([]uint, length)}
}

// A private function extends the size of bitset.
// The new BitSet object has a copy of the original
// and extends with zero values
func (b *BitSet) extends(length uint) {
	if b.bits == nil {
		b.bits = make([]uint, length)
	} else {
		newbits := make([]uint, length)
		copy(newbits, b.bits)
		b.bits = newbits
	}
	b.length = length
}

// Sets value of index to 1. If index is exceed the
// size of bitset, it extends the bitset. The position
// in bitset with value 1 indicates it is occupied.
func (b *BitSet) Set(i uint) {
	if i >= b.length {
		b.extends(i + 1)
	}
	b.bits[i] = 1
}

// Sets value of index to 0, indicates the position
// of bitset is released.
func (b *BitSet) Clear(i uint) {
	if i >= b.length {
		return
	}
	b.bits[i] = 0
}

// Increases value of bits i
func (b *BitSet) Add(i uint) {
	if i >= b.length {
		b.extends(i + 1)
	}
	b.bits[i] += 1
}

// Decreases value of bits i. Just return if value of
// bits i is zero.
func (b *BitSet) Sub(i uint) {
	if i >= b.length {
		return
	}
	if b.bits[i] > 0 {
		b.bits[i] -= 1
	}
}

// Clears all bits and sets them to 0
func (b *BitSet) ClearAll() {
	if b.bits != nil {
		for i := range b.bits {
			b.bits[i] = 0
		}
	}
}

// Checks whether the bits i is occupied. Returns
// true if is bits i is not zero, else false.
func (b *BitSet) Has(i uint) bool {
	if i >= b.length {
		return false
	}
	if b.bits[i] == 0 {
		return false
	}
	return true
}
