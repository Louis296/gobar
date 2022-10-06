package gobar

type BinaryIndexedTree []int

func NewBinaryIndexedTreeWithSize(size int) BinaryIndexedTree {
	return make(BinaryIndexedTree, size+1)
}

func NewBinaryIndexedTree(data []int) BinaryIndexedTree {
	ans := NewBinaryIndexedTreeWithSize(len(data))
	for i, num := range data {
		ans.Update(i+1, num)
	}
	return ans
}

func (t BinaryIndexedTree) Update(p, v int) {
	for ; p < len(t); p += lowBit(p) {
		t[p] += v
	}
}

func (t BinaryIndexedTree) Query(l, r int) int {
	if l >= len(t) {
		return 0
	}
	if r >= len(t) {
		r = len(t) - 1
	}
	return t.query(r) - t.query(l-1)
}

func (t BinaryIndexedTree) query(n int) int {
	ans := 0
	for p := n; p > 0; p -= lowBit(p) {
		ans += t[p]
	}
	return ans
}

func lowBit(x int) int {
	return x & (-x)
}
