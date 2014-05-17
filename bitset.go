package gobloom

type BitSet struct {
	length uint
	bits   []uint
}

func NewBitSet(length uint) *BitSet {
	return &BitSet{length, make([]uint, length)}
}

// extends length of bitset
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

// set value to 1
func (b *BitSet) Set(i uint) {
	if i >= b.length {
		b.extends(i + 1)
	}
	b.bits[i] = 1
}

// set value to 0
func (b *BitSet) Clear(i uint) {
	if i >= b.length {
		return
	}
	b.bits[i] = 0
}

// add 1
func (b *BitSet) Add(i uint) {
	if i >= b.length {
		b.extends(i + 1)
	}
	b.bits[i] += 1
}

// sub 1
func (b *BitSet) Sub(i uint) {
	if i >= b.length {
		return
	}
	if b.bits[i] > 0 {
		b.bits[i] -= 1
	}
}

// clear all bits
func (b *BitSet) ClearAll() {
	if b.bits != nil {
		for i := range b.bits {
			b.bits[i] = 0
		}
	}
}

func (b *BitSet) Has(i uint) bool {
	if i >= b.length {
		return false
	}
	if b.bits[i] == 0 {
		return false
	}
	return true
}
