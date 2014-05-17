package gobloom

import (
	"fmt"
	"testing"
)

func TestBF(t *testing.T) {
	bf := NewBloomFilter(100, 2)

	if bf.Contains([]byte("hello world")) == true {
		t.Error("did not work as expected")
	}

	bf.Add([]byte("yusong"))
	if bf.Contains([]byte("yusong")) == false {
		t.Error("did not work as expected")
	}

	bf.Add([]byte("yusonglam"))
	if bf.Contains([]byte("yusonglam")) == false {
		t.Error("did not work as expected")
	}

	fmt.Println(bf.EstimateCurrentFPR(), EstimateFPR(100, 5, 100))
}

func TestCBF(t *testing.T) {
	cbf := NewCountingBloomFilter(10, 3)

	if cbf.Contains([]byte("hello world")) == true {
		t.Error("did not work as expected")
	}

	cbf.Add([]byte("yusong"))
	if cbf.Contains([]byte("yusong")) == false {
		t.Error("did not work as expected")
	}

	cbf.Add([]byte("yusonglam"))
	if cbf.Contains([]byte("yusonglam")) == false {
		t.Error("did not work as expected")
	}

	fmt.Println(cbf.buckets)

	cbf.Remove([]byte("yusonglam"))
	if cbf.Contains([]byte("yusonglam")) == true {
		t.Error("did not work as expected")
	}

	fmt.Println(cbf.buckets)
}
