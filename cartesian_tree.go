package gobar

import "sort"

// CartesianTree is a tree that if you focus on k-val,
// it will look like a binary-search tree. Otherwise,
// if you focus on v-val, it is a min heap.
type CartesianTree struct {
	Root *cartesianTreeNode
}

type cartesianTreeNode struct {
	K     int
	V     int
	Left  *cartesianTreeNode
	Right *cartesianTreeNode
}

func NewCartesianTreeByIntList(arr []int) *CartesianTree {
	var kv [][]int
	for i, num := range arr {
		kv = append(kv, []int{i, num})
	}
	return NewCartesianTreeByKVList(kv)
}

func NewCartesianTreeByKVList(arr [][]int) *CartesianTree {
	if len(arr) == 0 {
		return &CartesianTree{Root: nil}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})
	var st []*cartesianTreeNode
	for i := range arr {
		now := &cartesianTreeNode{K: arr[i][0], V: arr[i][1]}
		for len(st) != 0 && st[len(st)-1].V > now.V {
			now.Left = st[len(st)-1]
			st = st[:len(st)-1]
		}
		if len(st) != 0 {
			st[len(st)-1].Right = now
		}
		st = append(st, now)
	}
	return &CartesianTree{Root: st[0]}
}
