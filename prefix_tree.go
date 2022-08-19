package gobar

type node struct {
	str      string
	indices  string
	children []*node
	value    interface{}
	priority uint32
}

type PrefixTree struct {
	root           *node
	size           int
	EnablePriority bool
}

func (t *PrefixTree) Insert(k string, v interface{}) bool {
	if t.root == nil {
		// init tree
		t.root = &node{
			str:      k,
			indices:  "",
			children: []*node{},
			value:    v,
			priority: 1,
		}
		t.size++
		return true
	}

	// insert key-value into tree
	p := t.root
	p.priority++
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
				priority: p.priority - 1,
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
					if t.EnablePriority {
						i = incrementChildPriority(p, i)
					}
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
			if t.EnablePriority {
				incrementChildPriority(p, len(p.children)-1)
			}
			t.size++
			return true
		}

		// insert key already in tree
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

		// now str is not key's prefix
		if l < len(p.str) {
			return nil, false
		}

		// now str is key's prefix
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

		// now str equals to key
		return p.value, true
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

func (t *PrefixTree) Size() int {
	return t.size
}

func incrementChildPriority(p *node, i int) int {
	p.children[i].priority++
	oldPos := i
	for ; i > 0; i-- {
		if p.children[i].priority > p.children[i-1].priority {
			p.children[i], p.children[i-1] = p.children[i-1], p.children[i]
		} else {
			break
		}
	}
	if i != oldPos {
		p.indices = p.indices[:i] + p.indices[oldPos:oldPos+1] + p.indices[i:oldPos] + p.indices[oldPos+1:]
	}

	return i
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
