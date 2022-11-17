package gobar

import (
	"testing"
)

func TestHeap(t *testing.T) {
	data := []int{2, 3, 1, 2, 30, 2, 9}
	h := NewHeap(data)
	var sorted []int
	for h.Size() != 0 {
		sorted = append(sorted, h.PopTop())
	}
	for i := 1; i < len(sorted); i++ {
		if sorted[i] < sorted[i-1] {
			t.FailNow()
		}
	}
}
