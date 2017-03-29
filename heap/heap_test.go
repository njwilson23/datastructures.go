package heap

import (
	"fmt"
	"testing"
)

func verifyMaxHeap(h *Heap) bool {
	for i := 0; i != h.size/2; i++ {
		if 2*(i+1)-1 < h.size && h.value[i] < h.value[2*(i+1)-1] {
			fmt.Printf("heap index %d (%.2f) not smaller than its left child (%.2f)\n", i, h.value[i], h.value[2*(i+1)-1])
			return false
		}
		if 2*(i+1) < h.size && h.value[i] < h.value[2*(i+1)] {
			fmt.Printf("heap index %d (%.2f) not smaller than its right child (%.2f)\n", i, h.value[i], h.value[2*(i+1)])
			return false
		}
	}
	return true
}

func TestMaximum(t *testing.T) {
	value := []float64{16, 4, 10, 14, 7, 9, 3, 2, 8, 1}
	label := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	h := BuildMaxHeap(value, label)

	l, v, err := h.Maximum()
	if err != nil {
		t.Error()
	}
	if v != 16 {
		t.Fail()
	}
	if l != 0 {
		t.Fail()
	}

	l, v, err = h.ExtractMaximum()
	if err != nil {
		t.Error()
	}
	if v != 16 {
		t.Fail()
	}
	if l != 0 {
		t.Fail()
	}
	if h.size != 9 {
		t.Fail()
	}

	l, v, err = h.Maximum()
	if err != nil {
		t.Error()
	}
	if v != 14 {
		t.Fail()
	}
	if l != 3 {
		t.Fail()
	}

}

func TestBuild(t *testing.T) {
	value := []float64{16, 4, 10, 14, 7, 9, 3, 2, 8, 1}
	label := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	h := BuildMaxHeap(value, label)
	if !verifyMaxHeap(h) {
		t.Fail()
	}
}
