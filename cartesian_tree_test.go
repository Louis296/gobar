package gobar

import (
	"fmt"
	"testing"
)

func TestCartesianTree(t *testing.T) {
	data := []int{1, 3, 2, 4, 5, 6}
	tree := NewCartesianTreeByIntList(data)
	fmt.Println(tree)
}
