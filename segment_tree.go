package gobar

type SegmentTree struct {
	tree  []int
	lazy  []int
	data  []int
	merge func(int, int) int
}

func NewSegmentTreeWithSize(size int, merge func(i, j int) int) SegmentTree {
	data := make([]int, size)
	return NewSegmentTree(data, merge)
}

func NewSegmentTree(data []int, merge func(i, j int) int) SegmentTree {
	n := len(data)
	data = append([]int{0}, data...)
	ans := SegmentTree{
		tree:  make([]int, 4*n+1),
		lazy:  make([]int, 4*n+1),
		data:  data,
		merge: merge,
	}
	ans.Build(1, n, 1)
	return ans
}

func (t SegmentTree) Build(l, r, k int) {
	if l == r {
		t.tree[k] = t.data[l]
		return
	}
	m := (l + r) >> 1
	t.Build(l, m, k<<1)
	t.Build(m+1, r, k<<1|1)
	t.pushUp(k)
}

func (t SegmentTree) pushUp(k int) {
	t.tree[k] = t.merge(t.tree[k<<1], t.tree[k<<1|1])
}

func (t SegmentTree) Query(l, r int) int {
	return t.query(l, r, 1, len(t.data)-1, 1)
}

func (t SegmentTree) query(L, R, l, r, k int) int {
	if L <= l && r <= R {
		return t.tree[k]
	}
	t.pushDown(k)
	m := (l + r) >> 1
	if m >= R {
		return t.query(L, R, l, m, k<<1)
	}
	if m <= L {
		return t.query(L, R, m+1, r, k<<1|1)
	}
	return t.merge(t.query(L, R, l, m, k<<1), t.query(L, R, m+1, r, k<<1|1))
}

// Update add data[pos] with v
func (t SegmentTree) Update(pos, v int) {
	t.update(pos, v, 1, len(t.data)-1, 1)
}

func (t SegmentTree) update(p, v, l, r, k int) {
	if l == r {
		t.tree[k] += v
		t.data[p] += v
		return
	}
	t.pushDown(k)
	m := (l + r) >> 1
	if p <= m {
		t.update(p, v, l, m, k<<1)
	} else {
		t.update(p, v, m+1, r, k<<1|1)
	}
	t.pushUp(k)
}

// UpdateSegment only can use when segment can be updated
func (t SegmentTree) UpdateSegment(l, r, v int) {
	t.updateSegment(l, r, v, 1, len(t.data)-1, 1)
}

func (t SegmentTree) pushDown(k int) {
	if t.lazy[k] != 0 {
		t.lazy[k<<1] += t.lazy[k]
		t.lazy[k<<1|1] += t.lazy[k]
		t.tree[k<<1] += t.lazy[k]
		t.tree[k<<1|1] += t.lazy[k]
		t.lazy[k] = 0
	}
}

func (t SegmentTree) updateSegment(L, R, v, l, r, k int) {
	if L <= l && r <= R {
		t.tree[k] += v
		t.lazy[k] += v
		return
	}
	t.pushDown(k)
	m := (l + r) >> 1
	if L <= m {
		t.updateSegment(L, R, v, l, m, k<<1)
	}
	if R > m {
		t.updateSegment(L, R, v, m+1, r, k<<1|1)
	}
	t.pushUp(k)
}
