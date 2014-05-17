package gobloom

import (
	"testing"
)

func TestBitSet(t *testing.T) {
	a := NewBitSet(10)

	a.Set(0)
	if a.bits[0] != 1 {
		t.Error("did not work as expected")
	}

	a.Add(15)
	a.Add(15)
	a.Add(15)
	if a.bits[15] != 3 {
		t.Error("did not work as expected")
	}

	a.Sub(15)
	if a.bits[15] != 2 {
		t.Error("did not work as expected")
	}

	a.Sub(3)
	if a.bits[3] != 0 {
		t.Error("did not work as expected")
	}

	a.ClearAll()
}
