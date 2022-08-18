package gobar

type node struct {
	str      string
	indices  string
	children []*node
	value    interface{}
}

type PrefixTree struct {
	root *node
	size int
}

func (t *PrefixTree) Insert(k string, v interface{}) bool {
	if t.root == nil {
		// init tree
		t.root = &node{
			str:      k,
			indices:  "",
			children: []*node{},
			value:    v,
		}
		return true
	}

	// insert key-value into tree
	p := t.root
walk:
	for {
		l := longestCommonPrefix(k, p.str)

		// split node
		if l < len(p.str) {
			newNode := &node{
				str:      p.str[l:],
				indices:  p.indices,
				children: p.children,
				value:    p.value,
			}
			p.str = p.str[:l]
			p.indices = string([]byte{newNode.str[0]})
			p.children = []*node{newNode}
			p.value = nil
		}

		if l < len(k) {
			k = k[l:]
			for i := range p.indices {
				if k[0] == p.indices[i] {
					p = p.children[i]
					continue walk
				}
			}
			newNode := &node{
				str:      k,
				indices:  "",
				children: []*node{},
				value:    v,
			}
			p.indices += k[0:1]
			p.children = append(p.children, newNode)
			t.size++
			return true
		}

		// node about key already in tree
		if p.value != nil {
			// duplicate key
			return false
		}
		p.value = v
		t.size++
		return true
	}
}

func (t *PrefixTree) Get(k string) (interface{}, bool) {
	if t.root == nil {
		return nil, false
	}

	p := t.root
walk:
	for {
		l := longestCommonPrefix(k, p.str)

		if l < len(p.str) {
			return nil, false
		}

		if l < len(k) {
			k = k[l:]
			for i := range p.indices {
				if k[0] == p.indices[i] {
					p = p.children[i]
					continue walk
				}
			}
			return nil, false
		}

		if p.str == k {
			return p.value, true
		}
		return nil, false
	}
}

func (t *PrefixTree) Delete(k string) bool {
	if t.root == nil {
		return false
	}
	var st []*node
	var indexSt []int
	p := t.root
walk:
	for {
		l := longestCommonPrefix(k, p.str)

		if l < len(p.str) {
			return false
		}

		if l < len(k) {
			k = k[l:]
			for i := range p.indices {
				if k[0] == p.indices[i] {
					st = append(st, p)
					indexSt = append(indexSt, i)
					p = p.children[i]
					continue walk
				}
			}
			return false
		}

		if p.str == k {
			p.value = nil
			break
		}
	}
	for len(st) != 0 && len(indexSt) != 0 {
		p = st[len(st)-1]
		st = st[:len(st)-1]

		i := indexSt[len(indexSt)-1]
		indexSt = indexSt[:len(indexSt)-1]

		if len(p.children[i].indices) == 0 && p.children[i].value == nil {
			p.children = append(p.children[:i], p.children[i+1:]...)
			p.indices = p.indices[:i] + p.indices[i+1:]
		} else {
			break
		}
	}
	if len(t.root.indices) == 0 {
		t.root = nil
	}
	t.size--
	return true
}

func longestCommonPrefix(a, b string) int {
	i := 0
	for ; i < min(len(a), len(b)); i++ {
		if a[i] != b[i] {
			break
		}
	}
	return i
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
