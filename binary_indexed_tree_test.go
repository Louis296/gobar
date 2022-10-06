package gobar

import (
	"fmt"
	"testing"
)

func TestBinaryIndexedTree(t *testing.T) {
	data := []int{1, 2, 3, 545, 12, 2, 3}
	tree := NewBinaryIndexedTree(data)
	fmt.Println(tree.Query(2, 3))
	fmt.Println(tree.Query(1, 100))

}
