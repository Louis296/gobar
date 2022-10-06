package gobar

import (
	"fmt"
	"testing"
)

func TestSegmentTree(t *testing.T) {
	data := []int{1, 2, 3, 45, 46, 2, 3, 4}
	tree := NewSegmentTree(data, func(i int, j int) int {
		return i + j
	})
	fmt.Println(tree.Query(5, 9))
	fmt.Println(tree.Query(1, 2))
	tree.Update(1, 2)
	fmt.Println(tree.Query(1, 2))

	fmt.Println(tree.Query(1, 10000000))
}

func TestNewSegmentTreeWithSize(t *testing.T) {
	tree := NewSegmentTreeWithSize(1000, func(i, j int) int {
		if i < j {
			return j
		}
		return i
	})
	fmt.Println(tree.Query(1, 100))
	tree.UpdateSegment(1, 100, 4)
	fmt.Println(tree.Query(1, 100))
	tree.UpdateSegment(2, 5, 3)
	fmt.Println(tree.Query(2, 3))

}
